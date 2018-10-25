package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStateChangeParams creates a new StateChangeParams object
// with the default values initialized.
func NewStateChangeParams() *StateChangeParams {
	var ()
	return &StateChangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStateChangeParamsWithTimeout creates a new StateChangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStateChangeParamsWithTimeout(timeout time.Duration) *StateChangeParams {
	var ()
	return &StateChangeParams{

		timeout: timeout,
	}
}

// NewStateChangeParamsWithContext creates a new StateChangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStateChangeParamsWithContext(ctx context.Context) *StateChangeParams {
	var ()
	return &StateChangeParams{

		Context: ctx,
	}
}

// NewStateChangeParamsWithHTTPClient creates a new StateChangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStateChangeParamsWithHTTPClient(client *http.Client) *StateChangeParams {
	var ()
	return &StateChangeParams{
		HTTPClient: client,
	}
}

/*StateChangeParams contains all the parameters to send to the API endpoint
for the state change operation typically these are written to a http.Request
*/
type StateChangeParams struct {

	/*OpID*/
	OpID *string
	/*Handle*/
	Handle string
	/*State*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the state change params
func (o *StateChangeParams) WithTimeout(timeout time.Duration) *StateChangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the state change params
func (o *StateChangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the state change params
func (o *StateChangeParams) WithContext(ctx context.Context) *StateChangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the state change params
func (o *StateChangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the state change params
func (o *StateChangeParams) WithHTTPClient(client *http.Client) *StateChangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the state change params
func (o *StateChangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the state change params
func (o *StateChangeParams) WithOpID(opID *string) *StateChangeParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the state change params
func (o *StateChangeParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithHandle adds the handle to the state change params
func (o *StateChangeParams) WithHandle(handle string) *StateChangeParams {
	o.SetHandle(handle)
	return o
}

// SetHandle adds the handle to the state change params
func (o *StateChangeParams) SetHandle(handle string) {
	o.Handle = handle
}

// WithState adds the state to the state change params
func (o *StateChangeParams) WithState(state string) *StateChangeParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the state change params
func (o *StateChangeParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *StateChangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param handle
	if err := r.SetPathParam("handle", o.Handle); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.State); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
