// Code generated by go-swagger; DO NOT EDIT.

package access_control

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
)

// NewDeleteInviteParams creates a new DeleteInviteParams object
// with the default values initialized.
func NewDeleteInviteParams() *DeleteInviteParams {
	var ()
	return &DeleteInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInviteParamsWithTimeout creates a new DeleteInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInviteParamsWithTimeout(timeout time.Duration) *DeleteInviteParams {
	var ()
	return &DeleteInviteParams{

		timeout: timeout,
	}
}

// NewDeleteInviteParamsWithContext creates a new DeleteInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInviteParamsWithContext(ctx context.Context) *DeleteInviteParams {
	var ()
	return &DeleteInviteParams{

		Context: ctx,
	}
}

// NewDeleteInviteParamsWithHTTPClient creates a new DeleteInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInviteParamsWithHTTPClient(client *http.Client) *DeleteInviteParams {
	var ()
	return &DeleteInviteParams{
		HTTPClient: client,
	}
}

/*DeleteInviteParams contains all the parameters to send to the API endpoint
for the delete invite operation typically these are written to a http.Request
*/
type DeleteInviteParams struct {

	/*InviteID
	  The invite ID to reference

	*/
	InviteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete invite params
func (o *DeleteInviteParams) WithTimeout(timeout time.Duration) *DeleteInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invite params
func (o *DeleteInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invite params
func (o *DeleteInviteParams) WithContext(ctx context.Context) *DeleteInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invite params
func (o *DeleteInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invite params
func (o *DeleteInviteParams) WithHTTPClient(client *http.Client) *DeleteInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invite params
func (o *DeleteInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInviteID adds the inviteID to the delete invite params
func (o *DeleteInviteParams) WithInviteID(inviteID string) *DeleteInviteParams {
	o.SetInviteID(inviteID)
	return o
}

// SetInviteID adds the inviteId to the delete invite params
func (o *DeleteInviteParams) SetInviteID(inviteID string) {
	o.InviteID = inviteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inviteId
	if err := r.SetPathParam("inviteId", o.InviteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
