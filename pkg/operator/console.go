package operator

import (
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"

	"github.com/openshift/console-operator/pkg/apis/console/v1alpha1"
)

func defaultConsole() *v1alpha1.Console {
	console := &v1alpha1.Console{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "console.openshift.io/v1alpha1",
			Kind:       "Console",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      OpenShiftConsoleName,
			Namespace: OpenShiftConsoleNamespace,
		},
		Spec: v1alpha1.ConsoleSpec{
			OperatorSpec: operatorv1alpha1.OperatorSpec{
				ManagementState: operatorv1alpha1.Managed,
			},
		},
	}
	console.SetDefaults()
	return console
}

func ApplyConsole() (*v1alpha1.Console, error) {
	console := defaultConsole()
	if err := sdk.Get(console); errors.IsNotFound(err) {
		if err = sdk.Create(console); err != nil {
			return nil, err
		}
	}
	return console, nil
}
