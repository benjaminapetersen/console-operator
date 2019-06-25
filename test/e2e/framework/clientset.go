package framework

import (
	"fmt"
	"testing"

	clientappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	clientcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"

	configv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	operatorclientv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	clientroutev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
)

// ClientSet is a set of Kubernetes clients.
type ClientSet struct {
	// embedded
	Core     clientcorev1.CoreV1Interface
	Apps     clientappsv1.AppsV1Interface
	Routes   clientroutev1.RouteV1Interface
	Operator operatorclientv1.ConsolesGetter
	Console  configv1.ConsolesGetter
}

// NewClientset creates a set of Kubernetes clients. The default kubeconfig is
// used if not provided.
func NewClientset(kubeconfig *restclient.Config) (*ClientSet, error) {
	var err error
	if kubeconfig == nil {
		kubeconfig, err = GetConfig()
		if err != nil {
			return nil, fmt.Errorf("unable to get kubeconfig: %s", err)
		}
	}

	clientset := &ClientSet{}
	clientset.Core, err = clientcorev1.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset.Apps, err = clientappsv1.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset.Routes, err = clientroutev1.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	operatorsClient, err := operatorclientv1.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset.Operator = operatorsClient

	return clientset, nil
}

// MustNewClientset is like NewClienset but aborts the test if clienset cannot
// be constructed.
func MustNewClientset(t *testing.T, kubeconfig *restclient.Config) *ClientSet {
	clientset, err := NewClientset(kubeconfig)
	if err != nil {
		t.Fatal(err)
	}
	return clientset
}
