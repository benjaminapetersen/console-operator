#!/usr/bin/env bash

YELLOW="\e[93m"
RESET="\e[97m"

# 1. delete operator deployment
echo -e "${YELLOW}deleting console-operator deployment...${RESET}"
oc delete deployment console-operator --namespace openshift-console-operator
oc delete configmap console-operator-lock -n openshift-console-operator

# 2. delete console namespace
echo -e "${YELLOW}deleting openshift-console namespace and all resources...${RESET}"
oc delete namespace openshift-console

# 3. recreate console namespace & rbac
echo -e "${YELLOW}creating openshift-console namespace & rbac roles...${RESET}"
oc apply -f manifests/02-namespace.yaml
oc apply -f manifests/03-rbac-role-cluster.yaml
oc apply -f manifests/03-rbac-role-ns-console.yaml
oc apply -f manifests/03-rbac-role-ns-openshift-config-managed.yaml
oc apply -f manifests/03-rbac-role-ns-operator.yaml
oc apply -f manifests/04-rbac-rolebinding.yaml

# 4. redeploy operator
echo -e "${YELLOW}redeploying console-operator...${RESET}"
oc create -f examples/07-operator-mine.yaml
echo -e "${YELLOW}checking pods...${RESET}"
oc get pods --namespace openshift-console-operator
oc get pods --namespace openshift-console
echo -e "${YELLOW}checking operator status...${RESET}"
oc get clusteroperator console
echo -e "${YELLOW}done. check console-operator logs${RESET}"

# test if all console resources are rolled out
