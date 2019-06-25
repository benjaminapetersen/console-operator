package e2e

import (
	"testing"

	"github.com/openshift/console-operator/test/e2e/framework"
)

func setupUnmanagedTestCase(t *testing.T) *framework.ClientSet {
	client := framework.MustNewClientset(t, nil)
	framework.MustUnmanageConsole(t, client)
	return client
}

func cleanUpUnmanagedTestCase(t *testing.T, client *framework.ClientSet) {
	framework.WaitForSettledState(t, client)
}

// TestUnmanaged() sets ManagementState:Unmanaged then deletes a set of console
// resources and verifies that the operator does not recreate them.
func TestUnmanaged(t *testing.T) {
	client := setupUnmanagedTestCase(t)
	defer framework.MustManageConsole(t, client)
	framework.DeleteAll(t, client)

	t.Logf("validating that the operator does not recreate deleted resources when ManagementState:Unmanaged...")
	err := framework.ConsoleResourcesUnavailable(client)
	if err != nil {
		t.Fatal(err)
	}
	cleanUpUnmanagedTestCase(t, client)
}

func TestEditUnmanagedConfigMap(t *testing.T) {
	client := setupUnmanagedTestCase(t)
	defer framework.MustManageConsole(t, client)

	err := patchAndCheckConfigMap(t, client, false)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	cleanUpUnmanagedTestCase(t, client)
}

func TestEditUnmanagedService(t *testing.T) {
	client := setupUnmanagedTestCase(t)
	defer framework.MustManageConsole(t, client)

	err := patchAndCheckService(t, client, false)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	cleanUpUnmanagedTestCase(t, client)
}

func TestEditUnmanagedRoute(t *testing.T) {
	client := setupUnmanagedTestCase(t)
	defer framework.MustManageConsole(t, client)

	err := patchAndCheckRoute(t, client, false)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	cleanUpUnmanagedTestCase(t, client)
}
