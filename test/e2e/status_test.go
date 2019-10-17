package e2e

import (
	"testing"

	operatorsv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/console-operator/test/e2e/framework"
)

func setupStatusTestCase(t *testing.T) (*framework.ClientSet, *operatorsv1.Console) {
	return framework.StandardSetup(t)
}

func cleanUpStatusTestCase(t *testing.T, client *framework.ClientSet) {
	framework.StandardCleanup(t, client)
}

// TODO: tests to write:
// - CustomBrand with incorrect keys
// - CustomBrand with missing source configmap
// -
func TestStatus(t *testing.T) {
	client, _ := setupStatusTestCase(t)
	defer cleanUpStatusTestCase(t, client)

}
