// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// BackupConfigGetOKCode is the HTTP code returned for type BackupConfigGetOK
const BackupConfigGetOKCode int = 200

/*BackupConfigGetOK return backup config

swagger:response backupConfigGetOK
*/
type BackupConfigGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.BackupConfig `json:"body,omitempty"`
}

// NewBackupConfigGetOK creates BackupConfigGetOK with default headers values
func NewBackupConfigGetOK() *BackupConfigGetOK {

	return &BackupConfigGetOK{}
}

// WithPayload adds the payload to the backup config get o k response
func (o *BackupConfigGetOK) WithPayload(payload *models.BackupConfig) *BackupConfigGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config get o k response
func (o *BackupConfigGetOK) SetPayload(payload *models.BackupConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupConfigGetBadRequestCode is the HTTP code returned for type BackupConfigGetBadRequest
const BackupConfigGetBadRequestCode int = 400

/*BackupConfigGetBadRequest invalid input, object invalid

swagger:response backupConfigGetBadRequest
*/
type BackupConfigGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigGetBadRequest creates BackupConfigGetBadRequest with default headers values
func NewBackupConfigGetBadRequest() *BackupConfigGetBadRequest {

	return &BackupConfigGetBadRequest{}
}

// WithPayload adds the payload to the backup config get bad request response
func (o *BackupConfigGetBadRequest) WithPayload(payload *models.Error) *BackupConfigGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config get bad request response
func (o *BackupConfigGetBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupConfigGetUnauthorizedCode is the HTTP code returned for type BackupConfigGetUnauthorized
const BackupConfigGetUnauthorizedCode int = 401

/*BackupConfigGetUnauthorized bad authentication

swagger:response backupConfigGetUnauthorized
*/
type BackupConfigGetUnauthorized struct {
}

// NewBackupConfigGetUnauthorized creates BackupConfigGetUnauthorized with default headers values
func NewBackupConfigGetUnauthorized() *BackupConfigGetUnauthorized {

	return &BackupConfigGetUnauthorized{}
}

// WriteResponse to the client
func (o *BackupConfigGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BackupConfigGetForbiddenCode is the HTTP code returned for type BackupConfigGetForbidden
const BackupConfigGetForbiddenCode int = 403

/*BackupConfigGetForbidden bad permissions

swagger:response backupConfigGetForbidden
*/
type BackupConfigGetForbidden struct {
}

// NewBackupConfigGetForbidden creates BackupConfigGetForbidden with default headers values
func NewBackupConfigGetForbidden() *BackupConfigGetForbidden {

	return &BackupConfigGetForbidden{}
}

// WriteResponse to the client
func (o *BackupConfigGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// BackupConfigGetNotFoundCode is the HTTP code returned for type BackupConfigGetNotFound
const BackupConfigGetNotFoundCode int = 404

/*BackupConfigGetNotFound item not found

swagger:response backupConfigGetNotFound
*/
type BackupConfigGetNotFound struct {
}

// NewBackupConfigGetNotFound creates BackupConfigGetNotFound with default headers values
func NewBackupConfigGetNotFound() *BackupConfigGetNotFound {

	return &BackupConfigGetNotFound{}
}

// WriteResponse to the client
func (o *BackupConfigGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// BackupConfigGetServiceUnavailableCode is the HTTP code returned for type BackupConfigGetServiceUnavailable
const BackupConfigGetServiceUnavailableCode int = 503

/*BackupConfigGetServiceUnavailable internal server error

swagger:response backupConfigGetServiceUnavailable
*/
type BackupConfigGetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigGetServiceUnavailable creates BackupConfigGetServiceUnavailable with default headers values
func NewBackupConfigGetServiceUnavailable() *BackupConfigGetServiceUnavailable {

	return &BackupConfigGetServiceUnavailable{}
}

// WithPayload adds the payload to the backup config get service unavailable response
func (o *BackupConfigGetServiceUnavailable) WithPayload(payload *models.Error) *BackupConfigGetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config get service unavailable response
func (o *BackupConfigGetServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigGetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
