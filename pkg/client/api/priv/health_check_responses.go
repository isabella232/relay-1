// Code generated by go-swagger; DO NOT EDIT.

package priv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// HealthCheckReader is a Reader for the HealthCheck structure.
type HealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthCheckOK creates a HealthCheckOK with default headers values
func NewHealthCheckOK() *HealthCheckOK {
	return &HealthCheckOK{}
}

/*HealthCheckOK handles this case with default header values.

Everything is healthy
*/
type HealthCheckOK struct {
	Payload *models.HealthResponse
}

func (o *HealthCheckOK) Error() string {
	return fmt.Sprintf("[GET /_priv/healthz][%d] healthCheckOK  %+v", 200, o.Payload)
}

func (o *HealthCheckOK) GetPayload() *models.HealthResponse {
	return o.Payload
}

func (o *HealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
