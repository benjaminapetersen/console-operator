package framework

import (
	"fmt"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"

	operatorsv1 "github.com/openshift/api/operator/v1"
	consoleapi "github.com/openshift/console-operator/pkg/api"
)

// set the operator config to a pristine state to start a next round of tests
// this should by default nullify out any customizations a user sets
func Pristine(t *testing.T, client *ClientSet) (*operatorsv1.Console, error) {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	copy := operatorConfig.DeepCopy()
	cleanSpec := operatorsv1.ConsoleSpec{}
	// we can set a default management state & log level, but
	// nothing else should be necessary
	cleanSpec.ManagementState = operatorsv1.Managed
	cleanSpec.LogLevel = operatorsv1.Normal
	copy.Spec = cleanSpec
	return client.Operator.Consoles().Update(copy)
}

func MustPristine(t *testing.T, client *ClientSet) *operatorsv1.Console {
	t.Helper()
	operatorConfig, err := Pristine(t, client)
	if err != nil {
		t.Fatal(err)
	}
	return operatorConfig
}

func isOperatorManaged(cr *operatorsv1.Console) bool {
	return cr.Spec.ManagementState == operatorsv1.Managed
}

func isOperatorUnmanaged(cr *operatorsv1.Console) bool {
	return cr.Spec.ManagementState == operatorsv1.Unmanaged
}

func isOperatorRemoved(cr *operatorsv1.Console) bool {
	return cr.Spec.ManagementState == operatorsv1.Removed
}

type operatorStateReactionFn func(cr *operatorsv1.Console) bool

func ensureConsoleIsInDesiredState(t *testing.T, client *ClientSet, state operatorsv1.ManagementState) error {
	t.Helper()
	var operatorConfig *operatorsv1.Console
	// var checkFunc func()
	var checkFunc operatorStateReactionFn

	switch state {
	case operatorsv1.Managed:
		checkFunc = isOperatorManaged
	case operatorsv1.Unmanaged:
		checkFunc = isOperatorUnmanaged
	case operatorsv1.Removed:
		checkFunc = isOperatorRemoved
	}

	err := wait.Poll(1*time.Second, AsyncOperationTimeout, func() (stop bool, err error) {
		operatorConfig, err = client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return checkFunc(operatorConfig), nil
	})
	if err != nil {
		DumpObject(t, "the latest observed state of the console resource", operatorConfig)
		DumpOperatorLogs(t, client)
		return fmt.Errorf("failed to wait to change console operator state to 'Removed': %s", err)
	}
	return nil
}

func ManageConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if isOperatorManaged(operatorConfig) {
		t.Logf("console operator already in 'Managed' state")
		return nil
	}

	t.Logf("changing console operator state to 'Managed'...")

	_, err = client.Operator.Consoles().Patch(consoleapi.ConfigResourceName, types.MergePatchType, []byte(`{"spec": {"managementState": "Managed"}}`))
	if err != nil {
		return err
	}
	if err := ensureConsoleIsInDesiredState(t, client, operatorsv1.Managed); err != nil {
		return fmt.Errorf("unable to change console operator state to 'Managed': %s", err)
	}

	err = ConsoleResourcesAvailable(client)
	if err != nil {
		t.Fatal(err)
	}

	return nil
}

func UnmanageConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if isOperatorUnmanaged(operatorConfig) {
		t.Logf("console operator already in 'Unmanaged' state")
		return nil
	}

	t.Logf("changing console operator state to 'Unmanaged'...")

	_, err = client.Operator.Consoles().Patch(consoleapi.ConfigResourceName, types.MergePatchType, []byte(`{"spec": {"managementState": "Unmanaged"}}`))
	if err != nil {
		return err
	}
	if err := ensureConsoleIsInDesiredState(t, client, operatorsv1.Unmanaged); err != nil {
		return fmt.Errorf("unable to change console operator state to 'Unmanaged': %s", err)
	}

	return nil
}

func RemoveConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if isOperatorRemoved(operatorConfig) {
		t.Logf("console operator already in 'Removed' state")
		return nil
	}

	t.Logf("changing console operator state to 'Removed'...")

	_, err = client.Operator.Consoles().Patch(consoleapi.ConfigResourceName, types.MergePatchType, []byte(`{"spec": {"managementState": "Removed"}}`))
	if err != nil {
		return err
	}
	if err := ensureConsoleIsInDesiredState(t, client, operatorsv1.Removed); err != nil {
		return fmt.Errorf("unable to change console operator state to 'Removed': %s", err)
	}

	return nil
}
func MustManageConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	if err := ManageConsole(t, client); err != nil {
		t.Fatal(err)
	}
	return nil
}

func MustUnmanageConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	if err := UnmanageConsole(t, client); err != nil {
		t.Fatal(err)
	}
	return nil
}

func MustRemoveConsole(t *testing.T, client *ClientSet) error {
	t.Helper()
	if err := RemoveConsole(t, client); err != nil {
		t.Fatal(err)
	}
	return nil
}

func MustNormalLogLevel(t *testing.T, client *ClientSet) error {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("checking if console operator LogLevel is set to 'Normal'...")
	if operatorConfig.Spec.LogLevel == operatorsv1.Normal {
		return nil
	}
	err = SetLogLevel(t, client, operatorsv1.Normal)
	if err != nil {
		t.Fatal(err)
	}
	return nil
}

func SetLogLevel(t *testing.T, client *ClientSet, logLevel operatorsv1.LogLevel) error {
	t.Helper()
	operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	deployment, err := GetConsoleDeployment(client)
	if err != nil {
		return err
	}
	currentDeploymentGeneration := deployment.ObjectMeta.Generation
	currentOperatorConfigGeneration := operatorConfig.ObjectMeta.Generation

	t.Logf("setting console operator to '%s' LogLevel ...", logLevel)
	_, err = client.Operator.Consoles().Patch(consoleapi.ConfigResourceName, types.MergePatchType, []byte(fmt.Sprintf(`{"spec": {"logLevel": "%s"}}`, logLevel)))
	if err != nil {
		return err
	}

	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		newOperatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
		newDeployment, err := GetConsoleDeployment(client)
		if err != nil {
			return false, nil
		}
		if GenerationChanged(newOperatorConfig.ObjectMeta.Generation, currentOperatorConfigGeneration) {
			return false, nil
		}
		if GenerationChanged(newDeployment.ObjectMeta.Generation, currentDeploymentGeneration) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GenerationChanged(oldGeneration, newGeneration int64) bool {
	return oldGeneration == newGeneration
}

// the operator is settled if conditions match:
// - available: true
// - progressing: false
// - degraded: false
func operatorIsSettled(operatorConfig *operatorsv1.Console) bool {
	settled := true
	for _, condition := range operatorConfig.Status.Conditions {
		if condition.Type == operatorsv1.OperatorStatusTypeAvailable {
			if condition.Status == operatorsv1.ConditionFalse {
				settled = false
				break
			}
		}
		if condition.Type == operatorsv1.OperatorStatusTypeProgressing {
			if condition.Status == operatorsv1.ConditionTrue {
				settled = false
				break
			}
		}
		if condition.Type == operatorsv1.OperatorStatusTypeDegraded {
			if condition.Status == operatorsv1.ConditionTrue {
				settled = false
				break
			}
		}
	}
	return settled
}

// A helper to ensure our operator config reaches a settled state before we
// begin the next test.
func WaitForSettledState(t *testing.T, client *ClientSet) (settled bool, err error) {
	t.Helper()
	fmt.Printf("waiting to reach settled state...\n")
	interval := 1 * time.Second
	// it should never take this long for a test to pass
	max := 240 * time.Second
	count := 0
	pollErr := wait.Poll(interval, max, func() (stop bool, err error) {
		// lets be informed about tests that take a long time to settle
		count++
		if count == 30 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		if count == 60 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		if count == 90 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		if count == 120 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		if count == 180 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		if count == 200 {
			fmt.Printf("waited %d seconds to reach settled state...\n", count)
		}
		operatorConfig, err := client.Operator.Consoles().Get(consoleapi.ConfigResourceName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		// first, wait until we are observing the correct generation
		if operatorConfig.Status.ObservedGeneration != operatorConfig.ObjectMeta.Generation {
			return false, nil
		}
		// then wait until the operator status settle
		return operatorIsSettled(operatorConfig), nil
	})
	if pollErr != nil {
		t.Errorf("operator has not reached settled state in %v attempts: %v", max, pollErr)
	}
	return true, nil

}
