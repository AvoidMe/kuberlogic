package app

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/util"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/generated/models"
	apiService "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/generated/restapi/operations/service"
	kuberlogiccomv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
)

// set this string to a required security grant for this action
const serviceDeleteSecGrant = "nonsense"

func (srv *Service) ServiceDeleteHandler(params apiService.ServiceDeleteParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	ns, name, err := util.SplitID(params.ServiceID)
	if err != nil {
		msg := "incorrect service id"
		srv.log.Errorw(msg, "error", err)
		return apiService.NewServiceDeleteBadRequest().WithPayload(&models.Error{
			Message: msg,
		})
	}

	r := new(kuberlogiccomv1alpha1.KuberLogicService)
	err = srv.kuberlogicClient.Get().
		Resource(serviceK8sResource).
		Namespace(ns).
		Name(name).
		Do(ctx).
		Into(r)
	if err != nil && util.ErrNotFound(err) {
		srv.log.Warnw("kuberlogic service not found",
			"namespace", ns, "name", name, "error", err)
		return apiService.NewServiceDeleteNotFound()
	} else if err != nil {
		return apiService.NewServiceDeleteServiceUnavailable().WithPayload(&models.Error{
			Message: "service not found",
		})
	}

	err = srv.kuberlogicClient.Delete().
		Resource(serviceK8sResource).
		Namespace(ns).
		Name(name).
		Do(ctx).
		Error()
	if err != nil {
		return apiService.NewServiceDeleteServiceUnavailable().WithPayload(&models.Error{
			Message: "error deleting service",
		})
	}

	return apiService.NewServiceDeleteOK()
}
