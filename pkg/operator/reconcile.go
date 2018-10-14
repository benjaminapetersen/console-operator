package operator

import (
	"github.com/openshift/console-operator/pkg/apis/console/v1alpha1"
)

func ReconcileConsole(cr *v1alpha1.Console) error {
	// TODO: scope the errors
	_, err := ApplyService(cr)
	if err != nil {
		return err
	}

	rt, err := ApplyRoute(cr)
	if err != nil {
		return err
	}

	_, err = ApplyConfigMap(cr, rt)
	if err != nil {
		return err
	}

	_, err = ApplyDeployment(cr)
	if err != nil {
		return err
	}

	_, _, err = ApplyOAuthClient(cr, rt)
	if err != nil {
		return err
	}

	return nil

}
