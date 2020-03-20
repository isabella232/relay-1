// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow runs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow runs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetWorkflowRun(params *GetWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunOK, error)

	GetWorkflowRunStepLog(params *GetWorkflowRunStepLogParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetWorkflowRunStepLogOK, *GetWorkflowRunStepLogPartialContent, error)

	GetWorkflowRuns(params *GetWorkflowRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunsOK, error)

	PatchWorkflowRun(params *PatchWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowRunOK, error)

	PatchWorkflowRunStep(params *PatchWorkflowRunStepParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowRunStepOK, error)

	RunWorkflow(params *RunWorkflowParams, authInfo runtime.ClientAuthInfoWriter) (*RunWorkflowCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetWorkflowRun gets a workflow run accessed with a workflow name and run number
*/
func (a *Client) GetWorkflowRun(params *GetWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRun",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflowName}/runs/{workflowRunNumber}",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkflowRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkflowRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWorkflowRunStepLog returns the log for a workflow step acessed by workflow name run number and step name
*/
func (a *Client) GetWorkflowRunStepLog(params *GetWorkflowRunStepLogParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetWorkflowRunStepLogOK, *GetWorkflowRunStepLogPartialContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunStepLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRunStepLog",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflowName}/runs/{workflowRunNumber}/steps/{workflowStepName}/logs",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkflowRunStepLogReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetWorkflowRunStepLogOK:
		return value, nil, nil
	case *GetWorkflowRunStepLogPartialContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowRunStepLogDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWorkflowRuns gets all the runs of a workflow
*/
func (a *Client) GetWorkflowRuns(params *GetWorkflowRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRuns",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflowName}/runs",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkflowRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkflowRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowRunsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchWorkflowRun updates properties of a workflow
*/
func (a *Client) PatchWorkflowRun(params *PatchWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchWorkflowRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchWorkflowRun",
		Method:             "PATCH",
		PathPattern:        "/api/workflows/{workflowName}/runs/{workflowRunNumber}",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchWorkflowRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchWorkflowRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchWorkflowRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchWorkflowRunStep updates properties of a workflow run step
*/
func (a *Client) PatchWorkflowRunStep(params *PatchWorkflowRunStepParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowRunStepOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchWorkflowRunStepParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchWorkflowRunStep",
		Method:             "PATCH",
		PathPattern:        "/api/workflows/{workflowName}/runs/{workflowRunNumber}/steps/{workflowStepName}",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchWorkflowRunStepReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchWorkflowRunStepOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchWorkflowRunStepDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RunWorkflow runs the given workflow
*/
func (a *Client) RunWorkflow(params *RunWorkflowParams, authInfo runtime.ClientAuthInfoWriter) (*RunWorkflowCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunWorkflowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "runWorkflow",
		Method:             "POST",
		PathPattern:        "/api/workflows/{workflowName}/runs",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20200131+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunWorkflowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunWorkflowCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunWorkflowDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
