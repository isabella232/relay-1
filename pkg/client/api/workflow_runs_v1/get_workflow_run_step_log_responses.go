// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWorkflowRunStepLogReader is a Reader for the GetWorkflowRunStepLog structure.
type GetWorkflowRunStepLogReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowRunStepLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowRunStepLogOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewGetWorkflowRunStepLogPartialContent(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWorkflowRunStepLogOK creates a GetWorkflowRunStepLogOK with default headers values
func NewGetWorkflowRunStepLogOK(writer io.Writer) *GetWorkflowRunStepLogOK {
	return &GetWorkflowRunStepLogOK{
		Payload: writer,
	}
}

/*GetWorkflowRunStepLogOK handles this case with default header values.

The full log
*/
type GetWorkflowRunStepLogOK struct {
	Payload io.Writer
}

func (o *GetWorkflowRunStepLogOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_name}/runs/{run_number}/steps/{step_name}/logs][%d] getWorkflowRunStepLogOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowRunStepLogOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetWorkflowRunStepLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRunStepLogPartialContent creates a GetWorkflowRunStepLogPartialContent with default headers values
func NewGetWorkflowRunStepLogPartialContent(writer io.Writer) *GetWorkflowRunStepLogPartialContent {
	return &GetWorkflowRunStepLogPartialContent{
		Payload: writer,
	}
}

/*GetWorkflowRunStepLogPartialContent handles this case with default header values.

The full (or partial) bytes of the log
*/
type GetWorkflowRunStepLogPartialContent struct {
	/*Returns the actual byte offsets being returned and how large the total log is, or "*" to indicate the log size is unknown since the step is not in a terminal state.
	 */
	ContentRange string

	Payload io.Writer
}

func (o *GetWorkflowRunStepLogPartialContent) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_name}/runs/{run_number}/steps/{step_name}/logs][%d] getWorkflowRunStepLogPartialContent  %+v", 206, o.Payload)
}

func (o *GetWorkflowRunStepLogPartialContent) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetWorkflowRunStepLogPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Range
	o.ContentRange = response.GetHeader("Content-Range")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
