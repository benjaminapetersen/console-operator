// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/console-operator/pkg/apis/console/v1"
	"github.com/openshift/console-operator/pkg/generated/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ConsoleV1Interface interface {
	RESTClient() rest.Interface
	ConsoleOperatorConfigsGetter
}

// ConsoleV1Client is used to interact with features provided by the console.openshift.io group.
type ConsoleV1Client struct {
	restClient rest.Interface
}

func (c *ConsoleV1Client) ConsoleOperatorConfigs() ConsoleOperatorConfigInterface {
	return newConsoleOperatorConfigs(c)
}

// NewForConfig creates a new ConsoleV1Client for the given config.
func NewForConfig(c *rest.Config) (*ConsoleV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ConsoleV1Client{client}, nil
}

// NewForConfigOrDie creates a new ConsoleV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ConsoleV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ConsoleV1Client for the given RESTClient.
func New(c rest.Interface) *ConsoleV1Client {
	return &ConsoleV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ConsoleV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
