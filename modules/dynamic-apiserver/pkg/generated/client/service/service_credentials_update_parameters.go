// Code generated by go-swagger; DO NOT EDIT.

package service

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

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// NewServiceCredentialsUpdateParams creates a new ServiceCredentialsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceCredentialsUpdateParams() *ServiceCredentialsUpdateParams {
	return &ServiceCredentialsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceCredentialsUpdateParamsWithTimeout creates a new ServiceCredentialsUpdateParams object
// with the ability to set a timeout on a request.
func NewServiceCredentialsUpdateParamsWithTimeout(timeout time.Duration) *ServiceCredentialsUpdateParams {
	return &ServiceCredentialsUpdateParams{
		timeout: timeout,
	}
}

// NewServiceCredentialsUpdateParamsWithContext creates a new ServiceCredentialsUpdateParams object
// with the ability to set a context for a request.
func NewServiceCredentialsUpdateParamsWithContext(ctx context.Context) *ServiceCredentialsUpdateParams {
	return &ServiceCredentialsUpdateParams{
		Context: ctx,
	}
}

// NewServiceCredentialsUpdateParamsWithHTTPClient creates a new ServiceCredentialsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceCredentialsUpdateParamsWithHTTPClient(client *http.Client) *ServiceCredentialsUpdateParams {
	return &ServiceCredentialsUpdateParams{
		HTTPClient: client,
	}
}

/* ServiceCredentialsUpdateParams contains all the parameters to send to the API endpoint
   for the service credentials update operation.

   Typically these are written to a http.Request.
*/
type ServiceCredentialsUpdateParams struct {

	/* ServiceCredentials.

	   service credentials
	*/
	ServiceCredentials models.ServiceCredentials

	/* ServiceID.

	   service Resource ID
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service credentials update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceCredentialsUpdateParams) WithDefaults() *ServiceCredentialsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service credentials update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceCredentialsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service credentials update params
func (o *ServiceCredentialsUpdateParams) WithTimeout(timeout time.Duration) *ServiceCredentialsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service credentials update params
func (o *ServiceCredentialsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service credentials update params
func (o *ServiceCredentialsUpdateParams) WithContext(ctx context.Context) *ServiceCredentialsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service credentials update params
func (o *ServiceCredentialsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service credentials update params
func (o *ServiceCredentialsUpdateParams) WithHTTPClient(client *http.Client) *ServiceCredentialsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service credentials update params
func (o *ServiceCredentialsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceCredentials adds the serviceCredentials to the service credentials update params
func (o *ServiceCredentialsUpdateParams) WithServiceCredentials(serviceCredentials models.ServiceCredentials) *ServiceCredentialsUpdateParams {
	o.SetServiceCredentials(serviceCredentials)
	return o
}

// SetServiceCredentials adds the serviceCredentials to the service credentials update params
func (o *ServiceCredentialsUpdateParams) SetServiceCredentials(serviceCredentials models.ServiceCredentials) {
	o.ServiceCredentials = serviceCredentials
}

// WithServiceID adds the serviceID to the service credentials update params
func (o *ServiceCredentialsUpdateParams) WithServiceID(serviceID string) *ServiceCredentialsUpdateParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service credentials update params
func (o *ServiceCredentialsUpdateParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceCredentialsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ServiceCredentials != nil {
		if err := r.SetBodyParam(o.ServiceCredentials); err != nil {
			return err
		}
	}

	// path param ServiceID
	if err := r.SetPathParam("ServiceID", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}