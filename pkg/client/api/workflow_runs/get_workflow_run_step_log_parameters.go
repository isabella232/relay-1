// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetWorkflowRunStepLogParams creates a new GetWorkflowRunStepLogParams object
// with the default values initialized.
func NewGetWorkflowRunStepLogParams() *GetWorkflowRunStepLogParams {
	var ()
	return &GetWorkflowRunStepLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowRunStepLogParamsWithTimeout creates a new GetWorkflowRunStepLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowRunStepLogParamsWithTimeout(timeout time.Duration) *GetWorkflowRunStepLogParams {
	var ()
	return &GetWorkflowRunStepLogParams{

		timeout: timeout,
	}
}

// NewGetWorkflowRunStepLogParamsWithContext creates a new GetWorkflowRunStepLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowRunStepLogParamsWithContext(ctx context.Context) *GetWorkflowRunStepLogParams {
	var ()
	return &GetWorkflowRunStepLogParams{

		Context: ctx,
	}
}

// NewGetWorkflowRunStepLogParamsWithHTTPClient creates a new GetWorkflowRunStepLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowRunStepLogParamsWithHTTPClient(client *http.Client) *GetWorkflowRunStepLogParams {
	var ()
	return &GetWorkflowRunStepLogParams{
		HTTPClient: client,
	}
}

/*GetWorkflowRunStepLogParams contains all the parameters to send to the API endpoint
for the get workflow run step log operation typically these are written to a http.Request
*/
type GetWorkflowRunStepLogParams struct {

	/*Range
	  Select a byte range, may specify an offset from the end (as in bytes=-200 to obtain the last 200 bytes).


	*/
	Range *string
	/*Follow
	  If true and the step is in progress, print known logs so far, then wait to send the next log chunk. Do not use in conjunction with the range header. Only use in conjunction with Accept: application/octet-stream.


	*/
	Follow *bool
	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string
	/*WorkflowRunNumber
	  Run number of the associated workflow

	*/
	WorkflowRunNumber int64
	/*WorkflowStepName
	  The name of the step in the associated workflow

	*/
	WorkflowStepName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithTimeout(timeout time.Duration) *GetWorkflowRunStepLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithContext(ctx context.Context) *GetWorkflowRunStepLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithHTTPClient(client *http.Client) *GetWorkflowRunStepLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRange adds the rangeVar to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithRange(rangeVar *string) *GetWorkflowRunStepLogParams {
	o.SetRange(rangeVar)
	return o
}

// SetRange adds the range to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetRange(rangeVar *string) {
	o.Range = rangeVar
}

// WithFollow adds the follow to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithFollow(follow *bool) *GetWorkflowRunStepLogParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithWorkflowName adds the workflowName to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithWorkflowName(workflowName string) *GetWorkflowRunStepLogParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WithWorkflowRunNumber adds the workflowRunNumber to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithWorkflowRunNumber(workflowRunNumber int64) *GetWorkflowRunStepLogParams {
	o.SetWorkflowRunNumber(workflowRunNumber)
	return o
}

// SetWorkflowRunNumber adds the workflowRunNumber to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetWorkflowRunNumber(workflowRunNumber int64) {
	o.WorkflowRunNumber = workflowRunNumber
}

// WithWorkflowStepName adds the workflowStepName to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) WithWorkflowStepName(workflowStepName string) *GetWorkflowRunStepLogParams {
	o.SetWorkflowStepName(workflowStepName)
	return o
}

// SetWorkflowStepName adds the workflowStepName to the get workflow run step log params
func (o *GetWorkflowRunStepLogParams) SetWorkflowStepName(workflowStepName string) {
	o.WorkflowStepName = workflowStepName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowRunStepLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Range != nil {

		// header param Range
		if err := r.SetHeaderParam("Range", *o.Range); err != nil {
			return err
		}

	}

	if o.Follow != nil {

		// query param follow
		var qrFollow bool
		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {
			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}

	}

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	// path param workflowRunNumber
	if err := r.SetPathParam("workflowRunNumber", swag.FormatInt64(o.WorkflowRunNumber)); err != nil {
		return err
	}

	// path param workflowStepName
	if err := r.SetPathParam("workflowStepName", o.WorkflowStepName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
