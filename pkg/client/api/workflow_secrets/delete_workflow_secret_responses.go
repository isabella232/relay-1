// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// DeleteWorkflowSecretReader is a Reader for the DeleteWorkflowSecret structure.
type DeleteWorkflowSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkflowSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteWorkflowSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWorkflowSecretOK creates a DeleteWorkflowSecretOK with default headers values
func NewDeleteWorkflowSecretOK() *DeleteWorkflowSecretOK {
	return &DeleteWorkflowSecretOK{}
}

/*DeleteWorkflowSecretOK handles this case with default header values.

The resource was deleted
*/
type DeleteWorkflowSecretOK struct {
	Payload *DeleteWorkflowSecretOKBody
}

func (o *DeleteWorkflowSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /api/workflows/{workflowName}/secrets/{workflowSecretName}][%d] deleteWorkflowSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkflowSecretOK) GetPayload() *DeleteWorkflowSecretOKBody {
	return o.Payload
}

func (o *DeleteWorkflowSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteWorkflowSecretOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkflowSecretDefault creates a DeleteWorkflowSecretDefault with default headers values
func NewDeleteWorkflowSecretDefault(code int) *DeleteWorkflowSecretDefault {
	return &DeleteWorkflowSecretDefault{
		_statusCode: code,
	}
}

/*DeleteWorkflowSecretDefault handles this case with default header values.

An error occurred
*/
type DeleteWorkflowSecretDefault struct {
	_statusCode int

	Payload *DeleteWorkflowSecretDefaultBody
}

// Code gets the status code for the delete workflow secret default response
func (o *DeleteWorkflowSecretDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkflowSecretDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/workflows/{workflowName}/secrets/{workflowSecretName}][%d] deleteWorkflowSecret default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWorkflowSecretDefault) GetPayload() *DeleteWorkflowSecretDefaultBody {
	return o.Payload
}

func (o *DeleteWorkflowSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteWorkflowSecretDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteWorkflowSecretDefaultBody Error response
swagger:model DeleteWorkflowSecretDefaultBody
*/
type DeleteWorkflowSecretDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this delete workflow secret default body
func (o *DeleteWorkflowSecretDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteWorkflowSecretDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWorkflowSecret default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWorkflowSecretDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWorkflowSecretDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteWorkflowSecretDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWorkflowSecretOKBody Information about the deleted resource
swagger:model DeleteWorkflowSecretOKBody
*/
type DeleteWorkflowSecretOKBody struct {

	// Deleted resource id
	// Required: true
	ResourceID *string `json:"resource_id"`

	// Was the resource successfully deleted?
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this delete workflow secret o k body
func (o *DeleteWorkflowSecretOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteWorkflowSecretOKBody) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("deleteWorkflowSecretOK"+"."+"resource_id", "body", o.ResourceID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteWorkflowSecretOKBody) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("deleteWorkflowSecretOK"+"."+"success", "body", o.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWorkflowSecretOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWorkflowSecretOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteWorkflowSecretOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
