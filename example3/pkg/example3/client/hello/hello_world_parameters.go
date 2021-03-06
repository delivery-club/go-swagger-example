// Code generated by go-swagger; DO NOT EDIT.

package hello

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

// NewHelloWorldParams creates a new HelloWorldParams object
// with the default values initialized.
func NewHelloWorldParams() *HelloWorldParams {

	return &HelloWorldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelloWorldParamsWithTimeout creates a new HelloWorldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelloWorldParamsWithTimeout(timeout time.Duration) *HelloWorldParams {

	return &HelloWorldParams{

		timeout: timeout,
	}
}

// NewHelloWorldParamsWithContext creates a new HelloWorldParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelloWorldParamsWithContext(ctx context.Context) *HelloWorldParams {

	return &HelloWorldParams{

		Context: ctx,
	}
}

// NewHelloWorldParamsWithHTTPClient creates a new HelloWorldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelloWorldParamsWithHTTPClient(client *http.Client) *HelloWorldParams {

	return &HelloWorldParams{
		HTTPClient: client,
	}
}

/*HelloWorldParams contains all the parameters to send to the API endpoint
for the hello world operation typically these are written to a http.Request
*/
type HelloWorldParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hello world params
func (o *HelloWorldParams) WithTimeout(timeout time.Duration) *HelloWorldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hello world params
func (o *HelloWorldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hello world params
func (o *HelloWorldParams) WithContext(ctx context.Context) *HelloWorldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hello world params
func (o *HelloWorldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hello world params
func (o *HelloWorldParams) WithHTTPClient(client *http.Client) *HelloWorldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hello world params
func (o *HelloWorldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HelloWorldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
