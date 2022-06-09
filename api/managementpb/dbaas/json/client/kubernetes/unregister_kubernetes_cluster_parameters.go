// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewUnregisterKubernetesClusterParams creates a new UnregisterKubernetesClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnregisterKubernetesClusterParams() *UnregisterKubernetesClusterParams {
	return &UnregisterKubernetesClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnregisterKubernetesClusterParamsWithTimeout creates a new UnregisterKubernetesClusterParams object
// with the ability to set a timeout on a request.
func NewUnregisterKubernetesClusterParamsWithTimeout(timeout time.Duration) *UnregisterKubernetesClusterParams {
	return &UnregisterKubernetesClusterParams{
		timeout: timeout,
	}
}

// NewUnregisterKubernetesClusterParamsWithContext creates a new UnregisterKubernetesClusterParams object
// with the ability to set a context for a request.
func NewUnregisterKubernetesClusterParamsWithContext(ctx context.Context) *UnregisterKubernetesClusterParams {
	return &UnregisterKubernetesClusterParams{
		Context: ctx,
	}
}

// NewUnregisterKubernetesClusterParamsWithHTTPClient creates a new UnregisterKubernetesClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnregisterKubernetesClusterParamsWithHTTPClient(client *http.Client) *UnregisterKubernetesClusterParams {
	return &UnregisterKubernetesClusterParams{
		HTTPClient: client,
	}
}

/* UnregisterKubernetesClusterParams contains all the parameters to send to the API endpoint
   for the unregister kubernetes cluster operation.

   Typically these are written to a http.Request.
*/
type UnregisterKubernetesClusterParams struct {
	// Body.
	Body UnregisterKubernetesClusterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unregister kubernetes cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterKubernetesClusterParams) WithDefaults() *UnregisterKubernetesClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unregister kubernetes cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterKubernetesClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) WithTimeout(timeout time.Duration) *UnregisterKubernetesClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) WithContext(ctx context.Context) *UnregisterKubernetesClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) WithHTTPClient(client *http.Client) *UnregisterKubernetesClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) WithBody(body UnregisterKubernetesClusterBody) *UnregisterKubernetesClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unregister kubernetes cluster params
func (o *UnregisterKubernetesClusterParams) SetBody(body UnregisterKubernetesClusterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UnregisterKubernetesClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
