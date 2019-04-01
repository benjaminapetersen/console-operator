package operator

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	operatorclientv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
)

// TODO: this should be done!
//
// initialize with the clients, etc needed to do its job
//    txn = NewStatusRecorder(foo, bar)
// when outer func completes, ok, process out all the statuses and then make the request
//    defer txn.Commit() // yay! ensure this happens before func() returns.
// pass in the config that will be acted upon when .Commit() is called
//    txn.Start(operatorConfig.status)
// add statuses when needed, trusting the .Commit() will process them
//  out and handle aggregation appropriately.  It can also ensure decisions
//  about conflicting status reports are correctly handled.
//    txn.Add(someNewThing)  // add to array
//    txn.Add(someNewThing)  // add to array
//    txn.Add(someNewThing)  // add to array
//
// Failing:  if 5 things come in, 4 not fail, 1 fail, we are failing
// Available: if 5 things come in, 4 available, 1 not, we are not available
// Progressing: if 5 things come in, 4 are not progresssing, 1 is, we are progressing
// Upgradable: etc

type StatusRecorder struct {
	client operatorclientv1.ConsoleInterface
}

func (sr *StatusRecorder) Start(status operatorv1.ConsoleStatus) {

}

func (sr *StatusRecorder) Add(condition operatorv1.OperatorCondition) {}

func (sr *StatusRecorder) Commit() {
	sr.client.UpdateStatus()
}

func NewStatusRecorder(
	consoleOperatorClient operatorclientv1.ConsoleInterface,
) *StatusRecorder {
	return &StatusRecorder{
		client: consoleOperatorClient,
	}
}
