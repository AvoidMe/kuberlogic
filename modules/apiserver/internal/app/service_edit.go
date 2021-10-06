package app

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
	apiService "github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/restapi/operations/service"
)

func (srv *Service) ServiceEditHandler(params apiService.ServiceEditParams, principal *models.Principal) middleware.Responder {
	service, errUpdate := srv.serviceStore.UpdateService(params.ServiceItem, principal, params.HTTPRequest.Context())
	if errUpdate != nil {
		srv.log.Errorw("error updating service", "error", errUpdate.Err)
		if errUpdate.Client {
			return apiService.NewServiceEditBadRequest().WithPayload(&models.Error{Message: errUpdate.ClientMsg})
		} else {
			return apiService.NewServiceEditServiceUnavailable().WithPayload(&models.Error{Message: errUpdate.ClientMsg})
		}
	}
	return apiService.NewServiceEditOK().WithPayload(service)
}
