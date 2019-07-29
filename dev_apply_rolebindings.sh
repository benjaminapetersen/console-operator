#!/usr/bin/env bash

echo "applying roles & rolebindings..."
oc apply -f ./manifests/04-rbac-rolebinding.yaml
oc apply -f ./manifests/03-rbac-role-cluster.yaml
oc apply -f ./manifests/03-rbac-role-ns-console.yaml
oc apply -f ./manifests/03-rbac-role-ns-openshift-config-managed.yaml
oc apply -f ./manifests/03-rbac-role-ns-operator.yaml
