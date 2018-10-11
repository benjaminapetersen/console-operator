package operator

import (
	// TODO: use when swapping up to client from Handler
	// "k8s.io/client-go/kubernetes/typed/core/v1"
	"github.com/openshift/console-operator/pkg/apis/console/v1alpha1"
)


// operator.Reconcile(cr)
// so ignore "resource exists", GET the resource, diff against expected, if not, UPDATE resource, loop.
//   it shouldn't loop infintely, however, at some point it ought to idle if things aren't changing
//   (until the next watch event fires)
// process should:
//   - burst when it is first reconciling to get everything into correct state
//   - update & reconcile only when things change. if no monkey business, should be idle
//   - wake up every <resyncPeriod> in main.go and do a reconcile again, just as a check
//   - note that API calls are expensive, so don't make them without good reason
// reconcile ought to do the following:
//   create deployment if not exists
//   create service if not exists
//   create route if not exists
//   create configmap if not exists
// 		which will look something like this:
//        sdk.Get(the-client)
//        if !exists
//          sdk.Get(the-route)
//          addRouteHostIfWeGotIt(the-client)
//           sdk.Create(the-client)
//        else
//          sdk.Get(the-route)
//          addRouteHostIfWeGotIt(the-client)
//          sdk.Update(the-client)
//   create oauthclient-secret if not exists
// but also
//   sync random secret between oauthclient & oauthclient-secret
//   sync route.host between route, oauthclient.redirectURIs & configmap.baseAddress

// func ReconcileConsole(cr *v1alpha1.Console, secretClient v1.SecretInterface) error {
// TODO: use when swapping up to client from Handler
func ReconcileConsole(cr *v1alpha1.Console) error {
	_, err := CreateServiceIfNotPresent(cr)
	if err != nil {
		return err
	}

	rt, err := CreateRouteIfNotPresent(cr)
	if err != nil {
		return err
	}

	_, err = CreateConsoleConfigMapIfNotPresent(cr, rt)
	if err != nil {
		return err
	}

	_, err = CreateConsoleDeploymentIfNotPresent(cr)
	if err != nil {
		return err
	}

	// TODO: use when swapping up to clients from Handler
	// _, _, err = UpdateOauthClientIfNotInSync(cr, rt, secretClient)
	_, _, err = UpdateOauthClientIfNotInSync(cr, rt)
	if err != nil {
		return err
	}

	return nil
}
