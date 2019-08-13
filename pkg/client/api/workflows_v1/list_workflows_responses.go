// Code generated by go-swagger; DO NOT EDIT.

package workflows_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// ListWorkflowsReader is a Reader for the ListWorkflows structure.
type ListWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListWorkflowsOK creates a ListWorkflowsOK with default headers values
func NewListWorkflowsOK() *ListWorkflowsOK {
	return &ListWorkflowsOK{}
}

/*ListWorkflowsOK handles this case with default header values.

An array of workflows
*/
type ListWorkflowsOK struct {
	Payload *models.Workflows
}

func (o *ListWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] listWorkflowsOK  %+v", 200, o.Payload)
}

func (o *ListWorkflowsOK) GetPayload() *models.Workflows {
	return o.Payload
}

func (o *ListWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workflows)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
