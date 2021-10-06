// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// DatabaseDeleteHandlerFunc turns a function with the right signature into a database delete handler
type DatabaseDeleteHandlerFunc func(DatabaseDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DatabaseDeleteHandlerFunc) Handle(params DatabaseDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DatabaseDeleteHandler interface for that can handle valid database delete params
type DatabaseDeleteHandler interface {
	Handle(DatabaseDeleteParams, *models.Principal) middleware.Responder
}

// NewDatabaseDelete creates a new http.Handler for the database delete operation
func NewDatabaseDelete(ctx *middleware.Context, handler DatabaseDeleteHandler) *DatabaseDelete {
	return &DatabaseDelete{Context: ctx, Handler: handler}
}

/*DatabaseDelete swagger:route DELETE /services/{ServiceID}/databases/{Database}/ service databaseDelete

DatabaseDelete database delete API

*/
type DatabaseDelete struct {
	Context *middleware.Context
	Handler DatabaseDeleteHandler
}

func (o *DatabaseDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDatabaseDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
