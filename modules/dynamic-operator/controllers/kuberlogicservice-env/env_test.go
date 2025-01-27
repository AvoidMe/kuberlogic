package kuberlogicservice_env

import (
	"context"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	cfg2 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/cfg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("KuberlogicserviceEnv Manager", func() {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))

	b := fake.NewClientBuilder()

	Context("When creating Kuberlogicservice", func() {
		client := b.WithScheme(scheme).Build()
		cfg := &cfg2.Config{
			Namespace: "kuberlogic",
			SvcOpts: struct {
				TLSSecretName string `envconfig:"optional"`
			}{
				TLSSecretName: "",
			},
		}

		kls := &v1alpha1.KuberLogicService{
			ObjectMeta: metav1.ObjectMeta{
				Name: "demo",
			},
			Spec:   v1alpha1.KuberLogicServiceSpec{},
			Status: v1alpha1.KuberLogicServiceStatus{},
		}

		It("Should prepare environment", func() {
			envMgr := New(client, kls, cfg)
			err := envMgr.SetupEnv(context.TODO())
			Expect(err).Should(BeNil())
			Expect(envMgr.NamespaceName).Should(Equal(kls.Name))

			By("Checking created objects")
			ns := &v1.Namespace{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: envMgr.NamespaceName}, ns)
			Expect(err).Should(BeNil())
			Expect(ns.GetName()).Should(Equal(envMgr.NamespaceName))

			netpol := &v12.NetworkPolicyList{}
			err = client.List(context.TODO(), netpol)
			Expect(err).Should(BeNil())
			Expect(len(netpol.Items)).Should(Equal(1))
		})
	})
})
