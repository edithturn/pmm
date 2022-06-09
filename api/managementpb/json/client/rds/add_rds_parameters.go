// Code generated by go-swagger; DO NOT EDIT.

package rds

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

// NewAddRDSParams creates a new AddRDSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRDSParams() *AddRDSParams {
	return &AddRDSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRDSParamsWithTimeout creates a new AddRDSParams object
// with the ability to set a timeout on a request.
func NewAddRDSParamsWithTimeout(timeout time.Duration) *AddRDSParams {
	return &AddRDSParams{
		timeout: timeout,
	}
}

// NewAddRDSParamsWithContext creates a new AddRDSParams object
// with the ability to set a context for a request.
func NewAddRDSParamsWithContext(ctx context.Context) *AddRDSParams {
	return &AddRDSParams{
		Context: ctx,
	}
}

// NewAddRDSParamsWithHTTPClient creates a new AddRDSParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRDSParamsWithHTTPClient(client *http.Client) *AddRDSParams {
	return &AddRDSParams{
		HTTPClient: client,
	}
}

/* AddRDSParams contains all the parameters to send to the API endpoint
   for the add RDS operation.

   Typically these are written to a http.Request.
*/
type AddRDSParams struct {
	// Body.
	Body AddRDSBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add RDS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRDSParams) WithDefaults() *AddRDSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add RDS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRDSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add RDS params
func (o *AddRDSParams) WithTimeout(timeout time.Duration) *AddRDSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add RDS params
func (o *AddRDSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add RDS params
func (o *AddRDSParams) WithContext(ctx context.Context) *AddRDSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add RDS params
func (o *AddRDSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add RDS params
func (o *AddRDSParams) WithHTTPClient(client *http.Client) *AddRDSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add RDS params
func (o *AddRDSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add RDS params
func (o *AddRDSParams) WithBody(body AddRDSBody) *AddRDSParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add RDS params
func (o *AddRDSParams) SetBody(body AddRDSBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRDSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
