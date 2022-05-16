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

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/generated/models"
)

// NewServiceAddParams creates a new ServiceAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceAddParams() *ServiceAddParams {
	return &ServiceAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceAddParamsWithTimeout creates a new ServiceAddParams object
// with the ability to set a timeout on a request.
func NewServiceAddParamsWithTimeout(timeout time.Duration) *ServiceAddParams {
	return &ServiceAddParams{
		timeout: timeout,
	}
}

// NewServiceAddParamsWithContext creates a new ServiceAddParams object
// with the ability to set a context for a request.
func NewServiceAddParamsWithContext(ctx context.Context) *ServiceAddParams {
	return &ServiceAddParams{
		Context: ctx,
	}
}

// NewServiceAddParamsWithHTTPClient creates a new ServiceAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceAddParamsWithHTTPClient(client *http.Client) *ServiceAddParams {
	return &ServiceAddParams{
		HTTPClient: client,
	}
}

/* ServiceAddParams contains all the parameters to send to the API endpoint
   for the service add operation.

   Typically these are written to a http.Request.
*/
type ServiceAddParams struct {

	/* ServiceItem.

	   service item
	*/
	ServiceItem *models.Service

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceAddParams) WithDefaults() *ServiceAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service add params
func (o *ServiceAddParams) WithTimeout(timeout time.Duration) *ServiceAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service add params
func (o *ServiceAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service add params
func (o *ServiceAddParams) WithContext(ctx context.Context) *ServiceAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service add params
func (o *ServiceAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service add params
func (o *ServiceAddParams) WithHTTPClient(client *http.Client) *ServiceAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service add params
func (o *ServiceAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceItem adds the serviceItem to the service add params
func (o *ServiceAddParams) WithServiceItem(serviceItem *models.Service) *ServiceAddParams {
	o.SetServiceItem(serviceItem)
	return o
}

// SetServiceItem adds the serviceItem to the service add params
func (o *ServiceAddParams) SetServiceItem(serviceItem *models.Service) {
	o.ServiceItem = serviceItem
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ServiceItem != nil {
		if err := r.SetBodyParam(o.ServiceItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
