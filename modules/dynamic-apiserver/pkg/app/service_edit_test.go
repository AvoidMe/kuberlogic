package app

import (
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/config"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
	apiService "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/service"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/util"
	cloudlinuxv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"net/http"
	"testing"
)

func TestServiceEditNotFound(t *testing.T) {
	expectedObject := &cloudlinuxv1alpha1.KuberLogicService{
		ObjectMeta: metav1.ObjectMeta{
			Name: "one",
		},
		Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
			Type:     "postgresql",
			Replicas: 1,
		},
	}

	tc := createTestClient(expectedObject, 404, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	params := apiService.ServiceEditParams{
		HTTPRequest: &http.Request{},
		ServiceID:   "not-found-id",
		ServiceItem: &models.Service{
			ID:       util.StrAsPointer("one"),
			Type:     util.StrAsPointer("postgresql"),
			Replicas: util.Int64AsPointer(1),
			Status:   "Unknown",
		},
	}

	checkResponse(srv.ServiceEditHandler(params, nil), t, 404, &models.Error{
		Message: "kuberlogic service not found: not-found-id",
	})
	tc.handler.ValidateRequestCount(t, 1)
}

func TestServiceEditSuccess(t *testing.T) {
	expectedObject := &cloudlinuxv1alpha1.KuberLogicService{
		ObjectMeta: metav1.ObjectMeta{
			Name: "one",
		},
		Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
			Type:     "postgresql",
			Replicas: 1,
			Limits: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("2Gi"),
			},
		},
	}

	tc := createTestClient(expectedObject, 200, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	service := &models.Service{
		ID:       util.StrAsPointer("one"),
		Type:     util.StrAsPointer("postgresql"),
		Replicas: util.Int64AsPointer(1),
		Limits: &models.Limits{
			Storage: "2Gi",
		},
		Status: "Unknown",
	}

	params := apiService.ServiceEditParams{
		HTTPRequest: &http.Request{},
		ServiceID:   "one",
		ServiceItem: service,
	}

	checkResponse(srv.ServiceEditHandler(params, nil), t, 200, service)
	tc.handler.ValidateRequestCount(t, 1)
}

func TestServiceEditForbidSetSubscription(t *testing.T) {
	expectedObject := &cloudlinuxv1alpha1.KuberLogicService{
		ObjectMeta: metav1.ObjectMeta{
			Name: "one",
		},
		Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
			Type:     "postgresql",
			Replicas: 1,
			Limits: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("2Gi"),
			},
		},
	}

	tc := createTestClient(expectedObject, 404, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	params := apiService.ServiceEditParams{
		HTTPRequest: &http.Request{},
		ServiceID:   "not-found-id",
		ServiceItem: &models.Service{
			ID:       util.StrAsPointer("one"),
			Type:     util.StrAsPointer("postgresql"),
			Replicas: util.Int64AsPointer(1),
			Limits: &models.Limits{
				Storage: "2Gi",
			},
			Status:       "Unknown",
			Subscription: "some-kind-of-subscription-id",
		},
	}

	checkResponse(srv.ServiceEditHandler(params, nil), t, 400, &models.Error{
		Message: "subscription cannot be changed",
	})
	tc.handler.ValidateRequestCount(t, 0)
}
