package consoleserver

import (
	"testing"

	"github.com/openshift/console-operator/pkg/api"

	"github.com/go-test/deep"
)

// Tests that the builder will return a correctly structured
// Console Server Config struct when builder.Config() is called
func TestConsoleServerCLIConfigBuilder(t *testing.T) {
	tests := []struct {
		name   string
		input  func() Config
		output Config
	}{
		{
			name: "Config builder should return default config if given no inputs",
			input: func() Config {
				b := &ConsoleServerCLIConfigBuilder{}
				return Config()
			},
			output: Config{
				Kind:       "ConsoleConfig",
				APIVersion: "console.openshift.io/v1",
				ServingInfo: ServingInfo{
					BindAddress: "https://0.0.0.0:8443",
					CertFile:    certFilePath,
					KeyFile:     keyFilePath,
				},
				ClusterInfo: ClusterInfo{
					ConsoleBasePath: "",
				},
				Auth: Auth{
					ClientID:            api.OpenShiftConsoleName,
					ClientSecretFile:    clientSecretFilePath,
					OAuthEndpointCAFile: oauthEndpointCAFilePath,
				},
				Customization: Customization{},
				Providers:     Providers{},
			},
		}, {
			name: "Config builder should handle cluster info",
			input: func() Config {
				b := &ConsoleServerCLIConfigBuilder{}
				return Config()
			},
			output: Config{
				Kind:       "ConsoleConfig",
				APIVersion: "console.openshift.io/v1",
				ServingInfo: ServingInfo{
					BindAddress: "https://0.0.0.0:8443",
					CertFile:    certFilePath,
					KeyFile:     keyFilePath,
				},
				ClusterInfo: ClusterInfo{
					ConsoleBasePath:    "",
					ConsoleBaseAddress: "https://foobar.com/host",
					MasterPublicURL:    "https://foobar.com/api",
				},
				Auth: Auth{
					ClientID:            api.OpenShiftConsoleName,
					ClientSecretFile:    clientSecretFilePath,
					OAuthEndpointCAFile: oauthEndpointCAFilePath,
					LogoutRedirect:      "https://foobar.com/logout",
				},
				Customization: Customization{},
				Providers:     Providers{},
			},
		}, {
			name: "Config builder should handle StatuspageID",
			input: func() Config {
				b := &ConsoleServerCLIConfigBuilder{}
				return Config()
			},
			output: Config{
				Kind:       "ConsoleConfig",
				APIVersion: "console.openshift.io/v1",
				ServingInfo: ServingInfo{
					BindAddress: "https://0.0.0.0:8443",
					CertFile:    certFilePath,
					KeyFile:     keyFilePath,
				},
				ClusterInfo: ClusterInfo{
					ConsoleBasePath: "",
				},
				Auth: Auth{
					ClientID:            api.OpenShiftConsoleName,
					ClientSecretFile:    clientSecretFilePath,
					OAuthEndpointCAFile: oauthEndpointCAFilePath,
				},
				Customization: Customization{},
				Providers: Providers{
					StatuspageID: "status-12345",
				},
			},
		}, {
			name: "Config builder should handle all inputs",
			input: func() Config {
				b := &ConsoleServerCLIConfigBuilder{}
				StatusPageID("status-12345")
				return Config()
			},
			output: Config{
				Kind:       "ConsoleConfig",
				APIVersion: "console.openshift.io/v1",
				ServingInfo: ServingInfo{
					BindAddress: "https://0.0.0.0:8443",
					CertFile:    certFilePath,
					KeyFile:     keyFilePath,
				},
				ClusterInfo: ClusterInfo{
					ConsoleBasePath: "",
					MasterPublicURL: "https://foobar.com/api",
				},
				Auth: Auth{
					ClientID:            api.OpenShiftConsoleName,
					ClientSecretFile:    clientSecretFilePath,
					OAuthEndpointCAFile: oauthEndpointCAFilePath,
					LogoutRedirect:      "https://foobar.com/logout",
				},
				Customization: Customization{
					Branding:             "okd",
					DocumentationBaseURL: "https://foobar.com/docs",
				},
				Providers: Providers{
					StatuspageID: "status-12345",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := deep.Equal(tt.input(), tt.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

// Tests that the builder will return correctly formatted YAML representing a
// Console Server Config struct when builder.ConfigYAML() is called.
// This YAML should be an exact representation of the struct from builder.Config()
// and the output should make use of the YAML tags embedded in the structs
// in config.go
func TestConsoleServerCLIConfigBuilderYAML(t *testing.T) {
	tests := []struct {
		name  string
		input func() ([]byte, error)
		// tests the YAML conversion output of the configmap
		output string
	}{
		{
			name: "Config builder should return default config if given no inputs",
			input: func() ([]byte, error) {
				b := &ConsoleServerCLIConfigBuilder{}
				return ConfigYAML()
			},
			output: `apiVersion: console.openshift.io/v1
kind: ConsoleConfig
servingInfo:
  bindAddress: https://0.0.0.0:8443
  certFile: /var/serving-cert/tls.crt
  keyFile: /var/serving-cert/tls.key
clusterInfo: {}
auth:
  clientID: console
  clientSecretFile: /var/oauth-config/clientSecret
  oauthEndpointCAFile: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
customization: {}
providers: {}
`,
		},
		{
			name: "Config builder should handle cluster info",
			input: func() ([]byte, error) {
				b := &ConsoleServerCLIConfigBuilder{}
				return ConfigYAML()
			},
			output: `apiVersion: console.openshift.io/v1
kind: ConsoleConfig
servingInfo:
  bindAddress: https://0.0.0.0:8443
  certFile: /var/serving-cert/tls.crt
  keyFile: /var/serving-cert/tls.key
clusterInfo:
  consoleBaseAddress: https://foobar.com/host
  masterPublicURL: https://foobar.com/api
auth:
  clientID: console
  clientSecretFile: /var/oauth-config/clientSecret
  oauthEndpointCAFile: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
  logoutRedirect: https://foobar.com/logout
customization: {}
providers: {}
`,
		},
		{
			name: "Config builder should handle StatuspageID",
			input: func() ([]byte, error) {
				b := &ConsoleServerCLIConfigBuilder{}
				return ConfigYAML()
			},
			output: `apiVersion: console.openshift.io/v1
kind: ConsoleConfig
servingInfo:
  bindAddress: https://0.0.0.0:8443
  certFile: /var/serving-cert/tls.crt
  keyFile: /var/serving-cert/tls.key
clusterInfo: {}
auth:
  clientID: console
  clientSecretFile: /var/oauth-config/clientSecret
  oauthEndpointCAFile: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
customization: {}
providers:
  statuspageID: status-12345
`,
		},
		{
			name: "Config builder should handle all inputs",
			input: func() ([]byte, error) {
				b := &ConsoleServerCLIConfigBuilder{}
				StatusPageID("status-12345")
				return ConfigYAML()
			},
			output: `apiVersion: console.openshift.io/v1
kind: ConsoleConfig
servingInfo:
  bindAddress: https://0.0.0.0:8443
  certFile: /var/serving-cert/tls.crt
  keyFile: /var/serving-cert/tls.key
clusterInfo:
  masterPublicURL: https://foobar.com/api
auth:
  clientID: console
  clientSecretFile: /var/oauth-config/clientSecret
  oauthEndpointCAFile: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
  logoutRedirect: https://foobar.com/logout
customization:
  branding: okd
  documentationBaseURL: https://foobar.com/docs
providers:
  statuspageID: status-12345
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, _ := tt.input()
			if diff := deep.Equal(string(input), tt.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
