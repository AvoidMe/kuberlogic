/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	compose "github.com/compose-spec/compose-go/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	//"github.com/go-logr/logr"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-operator/plugin/commons"
	pluginCompose "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/plugins/docker-compose/plugin/compose"
	"go.uber.org/zap"
	"strings"
)

type dockerComposeService struct {
	logger *zap.SugaredLogger
	spec   *compose.Project
}

func (d *dockerComposeService) GetCredentialsMethod(req commons.PluginRequestCredentialsMethod) *commons.PluginResponseCredentialsMethod {
	composeObjects := pluginCompose.NewComposeModel(d.spec, d.logger)
	r, err := composeObjects.GetCredentialsMethod(&req)
	if r == nil {
		ret := &commons.PluginResponseCredentialsMethod{
			Err: "failed to get update credentials method",
		}

		if err != nil {
			ret.Err = err.Error()
		}
		return ret
	}
	return r
}

func (d *dockerComposeService) Convert(req commons.PluginRequest) *commons.PluginResponse {
	res := &commons.PluginResponse{}

	dcModel := pluginCompose.NewComposeModel(d.spec, d.logger)
	objects, err := dcModel.Reconcile(&req)
	if err != nil {
		d.logger.Error(err, "error reconciling cluster objects")
		res.Err = err.Error()
		return res
	}

	for _, item := range objects {
		for gvk, object := range item {
			// do not return objects with empty name
			if object.GetName() == "" {
				continue
			}

			_ = res.AddUnstructuredObject(object, gvk)
		}
	}

	res.Service = dcModel.AccessServiceName()
	res.Protocol = commons.HTTPproto
	return res
}

func (d *dockerComposeService) Status(req commons.PluginRequest) *commons.PluginResponseStatus {
	status := &commons.PluginResponseStatus{
		IsReady: false,
	}

	dcModel := pluginCompose.NewComposeModel(d.spec, d.logger)
	ready, err := dcModel.Ready(&req)
	if err != nil {
		d.logger.Error(err.Error(), "error checking for readiness")
		status.Err = err.Error()
	}
	status.IsReady = ready

	return status
}

func (d *dockerComposeService) Types() *commons.PluginResponse {
	res := &commons.PluginResponse{}

	dcModel := pluginCompose.NewComposeModel(d.spec, d.logger)
	types := dcModel.Types()

	// we need to filter duplicates first
	for _, item := range types {
		for gvk, object := range item {
			_ = res.AddUnstructuredObject(object, gvk)
		}
	}

	return res
}

func (d *dockerComposeService) Default() *commons.PluginResponseDefault {
	defaults := &commons.PluginResponseDefault{
		Replicas: 1,
	}
	err := defaults.SetLimits(&corev1.ResourceList{
		corev1.ResourceStorage: resource.MustParse("1Gi"),
	})
	if err != nil {
		defaults.Err = err.Error()
	}
	return defaults
}

func (d *dockerComposeService) ValidateCreate(req commons.PluginRequest) *commons.PluginResponseValidation {
	return &commons.PluginResponseValidation{
		Err: validateRequest(&req),
	}
}

func (d *dockerComposeService) ValidateUpdate(req commons.PluginRequest) *commons.PluginResponseValidation {
	return &commons.PluginResponseValidation{
		Err: validateRequest(&req),
	}
}

func (d *dockerComposeService) ValidateDelete(_ commons.PluginRequest) *commons.PluginResponseValidation {
	return &commons.PluginResponseValidation{}
}

func newDockerComposeServicePlugin(composeProject *compose.Project, logger *zap.SugaredLogger) *dockerComposeService {
	return &dockerComposeService{
		spec:   composeProject,
		logger: logger,
	}
}

func validateRequest(req *commons.PluginRequest) string {
	var validateErrors []string
	if req.Replicas != 1 {
		validateErrors = append(validateErrors, "only 1 replica can be set")
	}

	return strings.Join(validateErrors, ", ")
}
