// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeKubermaticV1
}

var projectsResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "projects"}

var projectsKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "Project"}

// Get takes name of the project, and returns the corresponding project object, and an error if there is any.
func (c *FakeProjects) Get(name string, options v1.GetOptions) (result *kubermaticv1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectsResource, name), &kubermaticv1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Project), err
}

// List takes label and field selectors, and returns the list of Projects that match those selectors.
func (c *FakeProjects) List(opts v1.ListOptions) (result *kubermaticv1.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectsResource, projectsKind, opts), &kubermaticv1.ProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ProjectList{ListMeta: obj.(*kubermaticv1.ProjectList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectsResource, opts))
}

// Create takes the representation of a project and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Create(project *kubermaticv1.Project) (result *kubermaticv1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectsResource, project), &kubermaticv1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Project), err
}

// Update takes the representation of a project and updates it. Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Update(project *kubermaticv1.Project) (result *kubermaticv1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectsResource, project), &kubermaticv1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Project), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjects) UpdateStatus(project *kubermaticv1.Project) (*kubermaticv1.Project, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectsResource, "status", project), &kubermaticv1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Project), err
}

// Delete takes name of the project and deletes it. Returns an error if one occurs.
func (c *FakeProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(projectsResource, name), &kubermaticv1.Project{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ProjectList{})
	return err
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubermaticv1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectsResource, name, pt, data, subresources...), &kubermaticv1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Project), err
}
