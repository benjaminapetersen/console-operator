package e2e

import (
	"testing"

	operatorsv1 "github.com/openshift/api/operator/v1"

	"github.com/openshift/console-operator/test/e2e/framework"
)

func setupRecoverTestCase(t *testing.T) (*framework.ClientSet, *operatorsv1.Console) {
	return framework.StandardSetup(t)
}

func cleanupRecoverTestCase(t *testing.T, client *framework.ClientSet) {
	framework.StandardCleanup(t, client)
}

func TestRecoverFromDeletedSecrets(t *testing.T) {
	client, _ := setupRecoverTestCase(t)
	defer cleanupRecoverTestCase(t, client)

	// TODO: delete all secrets in openshift-console && openshift-console-operator
	// and ensure that the operator & console both recover and function.
}
