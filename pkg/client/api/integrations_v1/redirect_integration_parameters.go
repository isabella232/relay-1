// Code generated by go-swagger; DO NOT EDIT.

package integrations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRedirectIntegrationParams creates a new RedirectIntegrationParams object
// with the default values initialized.
func NewRedirectIntegrationParams() *RedirectIntegrationParams {
	var ()
	return &RedirectIntegrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectIntegrationParamsWithTimeout creates a new RedirectIntegrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRedirectIntegrationParamsWithTimeout(timeout time.Duration) *RedirectIntegrationParams {
	var ()
	return &RedirectIntegrationParams{

		timeout: timeout,
	}
}

// NewRedirectIntegrationParamsWithContext creates a new RedirectIntegrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewRedirectIntegrationParamsWithContext(ctx context.Context) *RedirectIntegrationParams {
	var ()
	return &RedirectIntegrationParams{

		Context: ctx,
	}
}

// NewRedirectIntegrationParamsWithHTTPClient creates a new RedirectIntegrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRedirectIntegrationParamsWithHTTPClient(client *http.Client) *RedirectIntegrationParams {
	var ()
	return &RedirectIntegrationParams{
		HTTPClient: client,
	}
}

/*RedirectIntegrationParams contains all the parameters to send to the API endpoint
for the redirect integration operation typically these are written to a http.Request
*/
type RedirectIntegrationParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*ID
	  ID of the integration we want redirect information for

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the redirect integration params
func (o *RedirectIntegrationParams) WithTimeout(timeout time.Duration) *RedirectIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect integration params
func (o *RedirectIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect integration params
func (o *RedirectIntegrationParams) WithContext(ctx context.Context) *RedirectIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect integration params
func (o *RedirectIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect integration params
func (o *RedirectIntegrationParams) WithHTTPClient(client *http.Client) *RedirectIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect integration params
func (o *RedirectIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the redirect integration params
func (o *RedirectIntegrationParams) WithAccept(accept string) *RedirectIntegrationParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the redirect integration params
func (o *RedirectIntegrationParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithID adds the id to the redirect integration params
func (o *RedirectIntegrationParams) WithID(id string) *RedirectIntegrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the redirect integration params
func (o *RedirectIntegrationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
