// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListAWSVPCSReader is a Reader for the ListAWSVPCS structure.
type ListAWSVPCSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSVPCSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAWSVPCSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAWSVPCSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSVPCSOK creates a ListAWSVPCSOK with default headers values
func NewListAWSVPCSOK() *ListAWSVPCSOK {
	return &ListAWSVPCSOK{}
}

/*ListAWSVPCSOK handles this case with default header values.

AWSVPCList
*/
type ListAWSVPCSOK struct {
	Payload models.AWSVPCList
}

func (o *ListAWSVPCSOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/vpcs][%d] listAWSVPCSOK  %+v", 200, o.Payload)
}

func (o *ListAWSVPCSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSVPCSDefault creates a ListAWSVPCSDefault with default headers values
func NewListAWSVPCSDefault(code int) *ListAWSVPCSDefault {
	return &ListAWSVPCSDefault{
		_statusCode: code,
	}
}

/*ListAWSVPCSDefault handles this case with default header values.

errorResponse
*/
type ListAWSVPCSDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s v p c s default response
func (o *ListAWSVPCSDefault) Code() int {
	return o._statusCode
}

func (o *ListAWSVPCSDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/vpcs][%d] listAWSVPCS default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSVPCSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
