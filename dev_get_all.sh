#!/usr/bin/env bash

echo "gathering resources..."
echo "routes..."
oc get route -n openshift-console
echo "services..."
oc get service -n openshift-console
echo "configmaps..."
oc get configmap -n openshift-console
echo "secrets..."
oc get secret -n openshift-console
echo "deployments..."
oc get deployment -n openshift-console
