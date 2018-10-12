# Console Operator 

An operator for OpenShift Console built using the operator-sdk.

The console-operator installs and maintains the web console on a cluster.

## Deploying the console-operator 

```bash 
# create from the manifests dir:
oc create -f manifests/
# if you get ordering errors, for example, cannot create rolebindings 
# because the namespace does not yet exist, run this after:  
oc apply -f manifests/ 
# now create a console custom resource:
oc create -f examples/cr.yaml 
```

## Running the console-operator locally 

```bash 
oc create -f manifests/ 
oc apply -f manifests/   # if anything is not created, see above
oc delete -f manifests/02-operator.yaml // we don't want this
# now build & run the operator locally:
operator-sdk up local --namespace=openshift-console
# now create a console custom resource:
oc create -f examples/cr.yaml  
```
