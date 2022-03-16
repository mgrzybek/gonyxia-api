package backoffice

// TODO: https://github.com/kubernetes/client-go
// TODO: https://helm.sh/docs/topics/advanced/#go-sdk

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/mgrzybek/gonyxia-api/internal/core"

	log "github.com/sirupsen/logrus"
)

// Kubernetes implements the OrchestratorAdaptor interface
type Kubernetes struct {
	// TODO: refactor in order to use only dynamic.Interface
	clientset *kubernetes.Clientset
	client    dynamic.Interface
}

/*
 * Constructors
 */

// NewKubernetes implements the Kubernetes constructor
func NewKubernetes(configFilePath *string) (Kubernetes, error) {
	var config *rest.Config
	var err error

	if configFilePath == nil {
		config, err = newInClusterConfigKubernetes()
	} else {
		config, err = newOutClusterConfigKubernetes(configFilePath)
	}
	if err != nil {
		log.Error(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err.Error())
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Error(err.Error())
	}

	result := Kubernetes{
		clientset: clientset,
		client:    client,
	}

	return result, result.Health()
}

func newInClusterConfigKubernetes() (*rest.Config, error) {
	log.Trace("Create an in-cluster Kubernetes object")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err.Error())
	}
	return config, err
}

func newOutClusterConfigKubernetes(configFilePath *string) (*rest.Config, error) {
	log.Trace("Create an out-cluster Kubernetes object using file ", configFilePath)
	config, err := clientcmd.BuildConfigFromFlags("", *configFilePath)
	if err != nil {
		log.Error(err.Error())
	}
	return config, err
}

/*
 * Interface’s implementations
 */

// Health lists the running PODs in order to test the orchestrator
func (k Kubernetes) Health() error {
	_, err := k.clientset.CoreV1().Pods("*").List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("Kubernetes endpoint is healthy")
	return nil
}

// UserCreate returns TODO
func (k Kubernetes) UserCreate() error {
	return fmt.Errorf("Not implemented")
}

// UserDelete returns TODO
func (k Kubernetes) UserDelete() error {
	return fmt.Errorf("Not implemented")
}

// NamespaceCreate creates a namespace
func (k Kubernetes) NamespaceCreate(name string, quota core.Quota, owner string) error {
	return fmt.Errorf("Not implemented")
}

// NamespaceDelete deletes the given namespace
func (k Kubernetes) NamespaceDelete(name string) error {
	return fmt.Errorf("Not implemented")
}

// RoleCreate creates a role in the RBAC system
func (k Kubernetes) RoleCreate() error {
	return fmt.Errorf("Not implemented")
}

// RoleDelete removes a role in the RBAC system
func (k Kubernetes) RoleDelete() error {
	return fmt.Errorf("Not implemented")
}

// GetQuota provides the ResourceQuotas Spec from the API
// https://iximiuz.com/en/posts/kubernetes-api-go-types-and-common-machinery/
// https://pkg.go.dev/k8s.io/api/core/v1#ResourceQuotaSpec
func (k Kubernetes) GetQuota(namespaceID string) (core.Quota, error) {
	if len(namespaceID) == 0 {
		err := fmt.Errorf("namespaceID is empty")
		log.Error(err)
		return core.Quota{}, err
	}

	rq := k.clientset.CoreV1().ResourceQuotas(namespaceID)
	result, err := rq.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Error(err)
		return core.Quota{}, err
	}

	if result.Items == nil || len(result.Items) == 0 {
		err := fmt.Errorf("Namespace %s not found", namespaceID)
		log.Warn(err)
		return core.Quota{}, err
	}

	q := result.Items[0].Spec

	reqStorage := q.Hard["requests.storage"]
	reqMemory := q.Hard["requests.memory"]
	reqCPU := q.Hard["requests.cpu"]

	limMemory := q.Hard["limits.memory"]
	limCPU := q.Hard["limits.cpu"]

	cPods := q.Hard["count.pods"]

	return core.Quota{
		RequestsMemory:  reqMemory.String(),
		RequestsCPU:     reqCPU.String(),
		LimitsMemory:    limMemory.String(),
		LimitsCPU:       limCPU.String(),
		RequestsStorage: reqStorage.String(),
		CountPods:       cPods.Value(),
	}, err
}

// SetQuota sets (create or update) quotas to the given namespace
func (k Kubernetes) SetQuota(q core.Quota, namespaceID string) error {
	quotaRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "ResourceQuota",
	}
	quota := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ResourceQuota",
			"metadata": map[string]interface{}{
				"name": namespaceID + "-quota",
			},
			"spec": map[string]interface{}{
				"hard": map[string]interface{}{
					"RequestMemory":   q.RequestsMemory,
					"ResquestCPU":     q.RequestsCPU,
					"LimitsMemory":    q.LimitsMemory,
					"LimitsCPU":       q.LimitsCPU,
					"RequestsStorage": q.RequestsStorage,
					"CountPods":       q.CountPods,
				},
			},
		},
	}

	log.Debug("Setting quota…")
	result, err := k.client.Resource(quotaRes).Namespace(namespaceID).Update(
		context.TODO(),
		quota,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	log.Debug("Quota updated: ", result)

	return nil
}
