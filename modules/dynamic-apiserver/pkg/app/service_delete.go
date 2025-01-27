package app

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
	apiService "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/service"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/util"
	kuberlogiccomv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (srv *Service) ServiceDeleteHandler(params apiService.ServiceDeleteParams, _ *models.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	r := new(kuberlogiccomv1alpha1.KuberLogicService)
	err := srv.kuberlogicClient.Get().
		Resource(serviceK8sResource).
		Name(params.ServiceID).
		Do(ctx).
		Into(r)
	if err != nil && util.CheckStatus(err, v1.StatusReasonNotFound) {
		msg := fmt.Sprintf("kuberlogic service not found: %s", params.ServiceID)
		srv.log.Warnw(msg, "error", err)
		return apiService.NewServiceDeleteNotFound().WithPayload(&models.Error{
			Message: msg,
		})
	} else if err != nil {
		msg := "error finding service"
		srv.log.Errorw(msg, "error", err)
		return apiService.NewServiceDeleteServiceUnavailable().WithPayload(&models.Error{
			Message: msg,
		})
	}

	err = srv.kuberlogicClient.Delete().
		Resource(serviceK8sResource).
		Name(params.ServiceID).
		Do(ctx).
		Error()
	if err != nil {
		msg := "error deleting service"
		srv.log.Errorw(msg, "error", err)
		return apiService.NewServiceDeleteServiceUnavailable().WithPayload(&models.Error{
			Message: msg,
		})
	}

	return apiService.NewServiceDeleteOK()
}
