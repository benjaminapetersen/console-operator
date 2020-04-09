package api

const (
	TargetNamespace    = "openshift-console"
	ConfigResourceName = "cluster"
)

// consts to maintain existing names of various sub-resources
const (
	ClusterOperatorName                     = "console"
	OpenShiftConsoleName                    = "console"
	OpenShiftConsoleNamespace               = TargetNamespace
	OpenShiftConsoleOperatorNamespace       = "openshift-console-operator"
	OpenShiftConsoleOperator                = "console-operator"
	OpenShiftConsoleConfigMapName           = "console-config"
	OpenShiftConsolePublicConfigMapName     = "console-public"
	ServiceCAConfigMapName                  = "service-ca"
	DefaultIngressCertConfigMapName         = "default-ingress-cert"
	OpenShiftConsoleDeploymentName          = OpenShiftConsoleName
	OpenShiftConsoleServiceName             = OpenShiftConsoleName
	OpenShiftConsoleRouteName               = OpenShiftConsoleName
	OpenShiftConsoleDownloadsRouteName      = "downloads"
	OAuthClientName                         = OpenShiftConsoleName
	OpenShiftConfigManagedNamespace         = "openshift-config-managed"
	OpenShiftConfigNamespace                = "openshift-config"
	OpenShiftMonitoringConfig               = "monitoring-shared-config"
	OpenShiftLoggingConfig                  = "logging-shared-config"
	OpenShiftCustomLogoConfigMapName        = "custom-logo"
	TrustedCAConfigMapName                  = "trusted-ca-bundle"
	TrustedCABundleKey                      = "ca-bundle.crt"
	TrustedCABundleMountDir                 = "/etc/pki/ca-trust/extracted/pem"
	TrustedCABundleMountFile                = "tls-ca-bundle.pem"
	OCCLIDownloadsCustomResourceName        = "oc-cli-downloads"
	OCCLIDownloadsLicenseCustomResourceName = OCCLIDownloadsCustomResourceName + "-license"
	ODOCLIDownloadsCustomResourceName       = "odo-cli-downloads"
)
