package e2e

import (
	"crypto/tls"
	"io/ioutil"

	"net/http"
	"testing"
	"time"

	// rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/openshift/console-operator/pkg/api"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operatorsv1 "github.com/openshift/api/operator/v1"
	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/console-operator/test/e2e/framework"
)

func setupMetricsEndpointTestCase(t *testing.T) (*framework.ClientSet, *operatorsv1.Console) {
	client, _ := framework.StandardSetup(t)
	routeForTest := tempRouteForTesting()
	_, err := client.Routes.Routes(api.OpenShiftConsoleOperatorNamespace).Create(routeForTest)
	if err != nil && !errors.IsAlreadyExists(err) {
		t.Fatalf("error: %s", err)
	}
	return client, nil
}

func cleanUpMetricsEndpointTestCase(t *testing.T, client *framework.ClientSet) {
	routeForTest := tempRouteForTesting()
	err := client.Routes.Routes(api.OpenShiftConsoleOperatorNamespace).Delete(routeForTest.Name, &metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	framework.StandardCleanup(t, client)
}

func TestMetricsEndpoint(t *testing.T) {
	client, _ := setupMetricsEndpointTestCase(t)
	defer cleanUpMetricsEndpointTestCase(t, client)

	tempRoute := tempRouteForTesting()
	routeForMetrics := ""
	err := wait.Poll(1*time.Second, 30*time.Second, func() (stop bool, err error) {
		tempRoute, err := client.Routes.Routes(api.OpenShiftConsoleOperatorNamespace).Get(tempRoute.Name, metav1.GetOptions{})
		if err != nil {
			t.Fatalf("error: %s", err)
		}
		if len(tempRoute.Spec.Host) == 0 {
			return false, err
		}
		routeForMetrics = "https://" + tempRoute.Spec.Host + "/metrics"
		// need port?
		// tempRoute.Spec.Port.TargetPort not sure this is correct
		// 8443? hard-coded?
		t.Logf("route for metrics: %v", routeForMetrics)
		return true, nil
	})
	t.Logf("route to /metrics: (%v) \n", routeForMetrics)
	// ignore default self signed cert for testing
	insecureTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	insecureClient := &http.Client{
		Transport: insecureTransport,
	}
	// ha, don't want the console, want the operator's own route!
	resp, err := insecureClient.Get(routeForMetrics)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("error: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	respString := string(bytes)
	t.Log("Response of metrics endpoint:")
	t.Log(respString)

}

// our metrics endpoint should not have a route, but this makes it
// easier to access the pod http://localhost:8443/metrics endpoint
// from our test
//
func tempRouteForTesting() *routev1.Route {
	return &routev1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "metrics",
			Namespace: api.OpenShiftConsoleOperatorNamespace,
		},
		Spec: routev1.RouteSpec{
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: "metrics",
			},
			Port: &routev1.RoutePort{
				TargetPort: intstr.FromString("https"),
			},
			TLS: &routev1.TLSConfig{
				Termination:                   routev1.TLSTerminationReencrypt,
				InsecureEdgeTerminationPolicy: routev1.InsecureEdgeTerminationPolicyRedirect,
			},
			WildcardPolicy: routev1.WildcardPolicyNone,
		},
	}
}

// perhaps we need an additional clusterrole...
//func tempClusterRole() rbacv1.ClusterRole {
//	return rbacv1.ClusterRole{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "console-operator-metrics",
//		},
//		Rules: []rbacv1.PolicyRule{
//			{
//				Verbs: []string{"GET", "LIST", "WATCH"},
//				// APIGroups:       nil,
//				// Resources:       nil,
//				// ResourceNames:   nil,
//				NonResourceURLs: []string{"/metrics"},
//			},
//		},
//	}
//}
//
//// and rolebinding...
//func tempClusterRoleBinding() rbacv1.ClusterRoleBinding {
//	return rbacv1.ClusterRoleBinding{}
//}
