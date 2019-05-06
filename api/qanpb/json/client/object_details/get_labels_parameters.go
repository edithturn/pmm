// Code generated by go-swagger; DO NOT EDIT.

package object_details

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

// NewGetLabelsParams creates a new GetLabelsParams object
// with the default values initialized.
func NewGetLabelsParams() *GetLabelsParams {
	var ()
	return &GetLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsParamsWithTimeout creates a new GetLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLabelsParamsWithTimeout(timeout time.Duration) *GetLabelsParams {
	var ()
	return &GetLabelsParams{

		timeout: timeout,
	}
}

// NewGetLabelsParamsWithContext creates a new GetLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLabelsParamsWithContext(ctx context.Context) *GetLabelsParams {
	var ()
	return &GetLabelsParams{

		Context: ctx,
	}
}

// NewGetLabelsParamsWithHTTPClient creates a new GetLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLabelsParamsWithHTTPClient(client *http.Client) *GetLabelsParams {
	var ()
	return &GetLabelsParams{
		HTTPClient: client,
	}
}

/*GetLabelsParams contains all the parameters to send to the API endpoint
for the get labels operation typically these are written to a http.Request
*/
type GetLabelsParams struct {

	/*Body*/
	Body GetLabelsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) WithTimeout(timeout time.Duration) *GetLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels params
func (o *GetLabelsParams) WithContext(ctx context.Context) *GetLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels params
func (o *GetLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) WithHTTPClient(client *http.Client) *GetLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get labels params
func (o *GetLabelsParams) WithBody(body GetLabelsBody) *GetLabelsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get labels params
func (o *GetLabelsParams) SetBody(body GetLabelsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
