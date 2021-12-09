package backoffice

// TODO: https://github.com/kubernetes/client-go
// TODO: https://helm.sh/docs/topics/advanced/#go-sdk

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/mgrzybek/gonyxia-api/internal/core"

	log "github.com/sirupsen/logrus"
)

type Kubernetes struct {
	clientset *kubernetes.Clientset
}

/*
 * Constructors
 */

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

	result := Kubernetes{
		clientset: clientset,
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
	log.Trace("Create an out-cluster Kubernetes object")
	config, err := clientcmd.BuildConfigFromFlags("", *configFilePath)
	if err != nil {
		log.Error(err.Error())
	}
	return config, err
}

/*
 * Interface’s implementations
 */

func (k Kubernetes) Health() error {
	_, err := k.clientset.CoreV1().Pods(
		os.Getenv("KUBERNETES_NAMESPACE"),
	).List(
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

func (k Kubernetes) UserCreate() error {
	return fmt.Errorf("Not implemented")
}
func (k Kubernetes) UserDelete() error {
	return fmt.Errorf("Not implemented")
}

func (k Kubernetes) NamespaceCreate(name string, quota core.Quota, owner string) error {
	return fmt.Errorf("Not implemented")
}

func (k Kubernetes) NamespaceDelete(name string) error {
	return fmt.Errorf("Not implemented")
}

func (k Kubernetes) RoleCreate() error {
	return fmt.Errorf("Not implemented")
}

func (k Kubernetes) RoleDelete() error {
	return fmt.Errorf("Not implemented")
}
