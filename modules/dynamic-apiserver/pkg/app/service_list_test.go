package app

import (
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/config"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
	apiService "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/service"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/util"
	cloudlinuxv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"net/http"
	"testing"
)

func TestServiceListEmpty(t *testing.T) {
	expectedObjects := &cloudlinuxv1alpha1.KuberLogicServiceList{
		Items: []cloudlinuxv1alpha1.KuberLogicService{},
	}

	tc := createTestClient(expectedObjects, 200, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	params := apiService.ServiceListParams{
		HTTPRequest: &http.Request{},
	}

	checkResponse(srv.ServiceListHandler(params, nil), t, 200, models.Services{})
	tc.handler.ValidateRequestCount(t, 1)
}

func TestServiceListMany(t *testing.T) {
	expectedObjects := &cloudlinuxv1alpha1.KuberLogicServiceList{
		Items: []cloudlinuxv1alpha1.KuberLogicService{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "one",
				},
				Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
					Type:     "postgresql",
					Replicas: 1,
					Domain:   "example.com",
				},
				Status: cloudlinuxv1alpha1.KuberLogicServiceStatus{
					Phase: "Running",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "two",
				},
				Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
					Type:     "mysql",
					Domain:   "example.com",
					Replicas: 2,
				},
				Status: cloudlinuxv1alpha1.KuberLogicServiceStatus{
					Phase: "Failed",
				},
			},
		},
	}

	tc := createTestClient(expectedObjects, 200, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	services := models.Services{
		{
			ID:       util.StrAsPointer("one"),
			Type:     util.StrAsPointer("postgresql"),
			Replicas: util.Int64AsPointer(1),
			Status:   "Running",
			Domain:   "example.com",
		},
		{
			ID:       util.StrAsPointer("two"),
			Type:     util.StrAsPointer("mysql"),
			Replicas: util.Int64AsPointer(2),
			Status:   "Failed",
			Domain:   "example.com",
		},
	}

	params := apiService.ServiceListParams{
		HTTPRequest: &http.Request{},
	}

	checkResponse(srv.ServiceListHandler(params, nil), t, 200, services)
	tc.handler.ValidateRequestCount(t, 1)
}

func TestServiceListWithSubscriptionFilter(t *testing.T) {
	expectedObjects := &cloudlinuxv1alpha1.KuberLogicServiceList{
		Items: []cloudlinuxv1alpha1.KuberLogicService{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "one",
					Labels: map[string]string{
						"subscription-id": "some-kind-of-subscription-id",
					},
				},
				Spec: cloudlinuxv1alpha1.KuberLogicServiceSpec{
					Type:     "postgresql",
					Replicas: 1,
					Domain:   "example.com",
				},
				Status: cloudlinuxv1alpha1.KuberLogicServiceStatus{
					Phase: "Running",
				},
			},
		},
	}

	tc := createTestClient(expectedObjects, 200, t)
	defer tc.server.Close()

	srv := &Service{
		log:              &TestLog{t: t},
		clientset:        fake.NewSimpleClientset(),
		kuberlogicClient: tc.client,
		config: &config.Config{
			Domain: "example.com",
		},
	}

	services := models.Services{
		{
			ID:           util.StrAsPointer("one"),
			Type:         util.StrAsPointer("postgresql"),
			Replicas:     util.Int64AsPointer(1),
			Status:       "Running",
			Subscription: "some-kind-of-subscription-id",
			Domain:       "example.com",
		},
	}

	params := apiService.ServiceListParams{
		HTTPRequest: &http.Request{},
	}

	checkResponse(srv.ServiceListHandler(params, nil), t, 200, services)
	tc.handler.ValidateRequestCount(t, 1)
}
