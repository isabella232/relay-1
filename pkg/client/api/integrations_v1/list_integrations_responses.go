// Code generated by go-swagger; DO NOT EDIT.

package integrations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// ListIntegrationsReader is a Reader for the ListIntegrations structure.
type ListIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIntegrationsOK creates a ListIntegrationsOK with default headers values
func NewListIntegrationsOK() *ListIntegrationsOK {
	return &ListIntegrationsOK{}
}

/*ListIntegrationsOK handles this case with default header values.

An array of integrations
*/
type ListIntegrationsOK struct {
	Payload *models.Integrations
}

func (o *ListIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /api/integrations][%d] listIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListIntegrationsOK) GetPayload() *models.Integrations {
	return o.Payload
}

func (o *ListIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integrations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
