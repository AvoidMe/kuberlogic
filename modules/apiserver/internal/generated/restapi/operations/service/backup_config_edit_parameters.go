// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// NewBackupConfigEditParams creates a new BackupConfigEditParams object
// no default values defined in spec.
func NewBackupConfigEditParams() BackupConfigEditParams {

	return BackupConfigEditParams{}
}

// BackupConfigEditParams contains all the bound params for the backup config edit operation
// typically these are obtained from a http.Request
//
// swagger:parameters backupConfigEdit
type BackupConfigEditParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Backup credential item to edit
	  Required: true
	  In: body
	*/
	BackupConfig *models.BackupConfig
	/*service Resource ID
	  Required: true
	  Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?:[a-z0-9]([-a-z0-9]*[a-z0-9])?
	  In: path
	*/
	ServiceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBackupConfigEditParams() beforehand.
func (o *BackupConfigEditParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.BackupConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("backupConfig", "body", ""))
			} else {
				res = append(res, errors.NewParseError("backupConfig", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.BackupConfig = &body
			}
		}
	} else {
		res = append(res, errors.Required("backupConfig", "body", ""))
	}
	rServiceID, rhkServiceID, _ := route.Params.GetOK("ServiceID")
	if err := o.bindServiceID(rServiceID, rhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindServiceID binds and validates parameter ServiceID from path.
func (o *BackupConfigEditParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceID = raw

	if err := o.validateServiceID(formats); err != nil {
		return err
	}

	return nil
}

// validateServiceID carries on validations for parameter ServiceID
func (o *BackupConfigEditParams) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Pattern("ServiceID", "path", o.ServiceID, `[a-z0-9]([-a-z0-9]*[a-z0-9])?:[a-z0-9]([-a-z0-9]*[a-z0-9])?`); err != nil {
		return err
	}

	return nil
}