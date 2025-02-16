/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusterdeletion

import (
	"strings"

	"k8c.io/reconciler/pkg/reconciling"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	utilpointer "k8s.io/utils/pointer"
)

// creationPreventingWebhook returns a ValidatingWebhookConfiguration that is intentionally defunct
// and will prevent all creation requests from succeeding.
func creationPreventingWebhook(apiGroup string, resources []string) reconciling.NamedValidatingWebhookConfigurationReconcilerFactory {
	failurePolicy := admissionregistrationv1.Fail
	sideEffects := admissionregistrationv1.SideEffectClassNone

	return func() (string, reconciling.ValidatingWebhookConfigurationReconciler) {
		return "kubernetes-cluster-cleanup-" + strings.Join(resources, "-"),
			func(vwc *admissionregistrationv1.ValidatingWebhookConfiguration) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
				if vwc.Annotations == nil {
					vwc.Annotations = map[string]string{}
				}
				vwc.Annotations[annotationKeyDescription] = "This webhook configuration exists to prevent creation of any new stateful resources in a cluster that is currently being terminated"

				// This only gets set when the APIServer supports it, so carry it over
				var scope *admissionregistrationv1.ScopeType
				if len(vwc.Webhooks) != 1 {
					vwc.Webhooks = []admissionregistrationv1.ValidatingWebhook{{}}
				} else if len(vwc.Webhooks[0].Rules) > 0 {
					scope = vwc.Webhooks[0].Rules[0].Scope
				}
				// Must be a domain with at least three segments separated by dots
				vwc.Webhooks[0].Name = "kubernetes.cluster.cleanup"
				vwc.Webhooks[0].ClientConfig = admissionregistrationv1.WebhookClientConfig{
					URL: utilpointer.String("https://127.0.0.1:1"),
				}
				vwc.Webhooks[0].Rules = []admissionregistrationv1.RuleWithOperations{
					{
						Operations: []admissionregistrationv1.OperationType{admissionregistrationv1.Create},
						Rule: admissionregistrationv1.Rule{
							APIGroups:   []string{apiGroup},
							APIVersions: []string{"*"},
							Resources:   resources,
							Scope:       scope,
						},
					},
				}
				vwc.Webhooks[0].FailurePolicy = &failurePolicy
				vwc.Webhooks[0].SideEffects = &sideEffects
				vwc.Webhooks[0].AdmissionReviewVersions = []string{"v1"}
				return vwc, nil
			}
	}
}
