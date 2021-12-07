package backoffice

// TODO: https://github.com/kubernetes/client-go
// TODO: https://helm.sh/docs/topics/advanced/#go-sdk

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/mgrzybek/gonyxia-api/internal/core"

	log "github.com/sirupsen/logrus"
)

type Kubernetes struct {
	clientset *kubernetes.Clientset
}

/*
 * Constructors
 */

func NewKubernetes(user, password, token *string) Kubernetes {
	if ( user == nil || password == nil ) && token == nil {
		return newInClusterKubernetes()
	}
	return newOutClusterKubernetes(user, password, token)
}

func newInClusterKubernetes() Kubernetes {
	log.Trace("Create an in-cluster Kubernetes object")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err.Error())
	}

	result := Kubernetes{
		clientset: clientset,
	}
	result.Health()

	return result
}

func newOutClusterKubernetes(user, password, token *string) Kubernetes {
	log.Trace("Create an out-cluster Kubernetes object")
	return Kubernetes{}
}

/*
 * Interface’s implementations
 */

func (k Kubernetes) Health() error {
	_, err := k.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (k Kubernetes) UserCreate() error {
	return fmt.Errorf("Not implemented")
}
func (k Kubernetes) UserDelete() error	{
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