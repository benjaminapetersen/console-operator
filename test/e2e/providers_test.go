package e2e

import (
	"strings"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	operatorsv1 "github.com/openshift/api/operator/v1"
	consoleapi "github.com/openshift/console-operator/pkg/api"
	"github.com/openshift/console-operator/test/e2e/framework"
)

const (
	statuspageIDField = "statuspageID"
	providersField    = "providers"
)

func setupProvidersTestCase(t *testing.T) (*framework.ClientSet, *operatorsv1.Console) {
	client, operatorConfig := framework.StandardSetup(t)
	return client, operatorConfig
}

func cleanupProvidersTestCase(t *testing.T, client *framework.ClientSet) {
	framework.StandardCleanup(t, client)
}

func TestProvidersSetStatuspageID(t *testing.T) {
	client, _ := setupProvidersTestCase(t)
	defer cleanupProvidersTestCase(t, client)
	expectedStatuspageID := "id-1234"
	currentStatuspageID := ""
	setOperatorConfigStatuspageIDProvider(t, client, expectedStatuspageID)

	err := wait.Poll(1*time.Second, pollTimeout, func() (stop bool, err error) {
		currentStatuspageID = getConsoleProviderField(t, client, statuspageIDField)
		if expectedStatuspageID == currentStatuspageID {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		t.Errorf("error: expected '%s' statuspageID, got '%s': '%v'", expectedStatuspageID, currentStatuspageID, err)
	}
}

func TestProvidersSetStatuspageIDEmpty(t *testing.T) {
	client, _ := setupProvidersTestCase(t)
	defer cleanupProvidersTestCase(t, client)
	statuspageID := ""
	currentProviders := ""
	expectedProviders := "{}"
	setOperatorConfigStatuspageIDProvider(t, client, statuspageID)

	err := wait.Poll(1*time.Second, pollTimeout, func() (stop bool, err error) {
		currentProviders = getConsoleProviderField(t, client, providersField)
		if currentProviders == expectedProviders {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		t.Errorf("error: expected '%s' statuspageID, got '%s': '%v'", expectedProviders, currentProviders, err)
	}
}

func getConsoleProviderField(t *testing.T, client *framework.ClientSet, providerField string) string {
	cm, err := framework.GetConsoleConfigMap(client)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	data := cm.Data["console-config.yaml"]
	field := ""
	temp := strings.Split(data, "\n")
	for _, item := range temp {
		if strings.Contains(item, providerField) {
			field = strings.TrimSpace(strings.Split(item, ":")[1])
			break
		}
	}
	return field
}

func setOperatorConfigStatuspageIDProvider(t *testing.T, client *framework.ClientSet, statuspageID string) {
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("could not get operator config, %v", err)
	}
	t.Logf("setting statuspageID to '%s'", statuspageID)
	operatorConfig.Spec = operatorsv1.ConsoleSpec{
		OperatorSpec: operatorsv1.OperatorSpec{
			ManagementState: "Managed",
		},
		Providers: operatorsv1.ConsoleProviders{
			Statuspage: &operatorsv1.StatuspageProvider{
				PageID: statuspageID,
			},
		},
	}
	_, err = client.Operator.Consoles().Update(operatorConfig)
	if err != nil {
		t.Fatalf("could not update operator config providers statupageID: %v", err)
	}
}
