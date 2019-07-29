#!/usr/bin/env bash

# after running ./dev_cluster_install.sh, run this script to
# configure the cluster for development purposes. Essentially,
# disable CVO, eliminate some noise,set replicas to 1, etc.

YELLOW="\e[93m"
RESET="\e[97m"

echo -e "${YELLOW}switch to project openshift-console-operator${RESET}"
oc project openshift-console-operator

echo -e "${YELLOW}scaling down CVO...${RESET}"
oc scale deployment cluster-version-operator --replicas 0 --namespace openshift-cluster-version
# oc scale deployment cluster-version-operator --replicas 1 --namespace openshift-cluster-version
oc get deployment cluster-version-operator --namespace openshift-cluster-version

echo -e "${YELLOW}scaling down default console-operator...${RESET}"
oc scale deployment console-operator --replicas 0 --namespace openshift-console-operator
oc get deployment console-operator --namespace openshift-console-operator

echo -e "${YELLOW}deploying alternative operator...${RESET}"
oc apply -f examples/07-operator-mine.yaml
oc apply -f examples/07-downloads-deployment-mine.yaml
oc delete configmap console-operator-lock
oc get deployment console-operator --namespace openshift-console-operator

echo -e "${YELLOW}cycling console deployments...${RESET}"
oc delete deployment console --namespace openshift-console
oc delete deployment downloads --namespace openshift-console
sleep 4
oc get deployment console --namespace openshift-console

# echo -e "${YELLOW}creating branding configmap...${RESET}"
# oc create -f examples/origin-branding-configmap-online.yaml

# deploy kube dashboard, why not?
# most of what is created goes in kube-system namespace
echo -e "${YELLOW}deploying kube dashboard...${RESET}"
oc apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

echo -e ""
echo -e "${YELLOW}given the above success:${RESET}"
echo -e "${YELLOW}CVO is no longer managing the cluster${RESET}"
echo -e "${YELLOW}console deployment deleted${RESET}"
echo -e "${YELLOW}downloads deployment deleted${RESET}"
echo -e "${YELLOW}Now, rebuild your operator image, push to your image repository, and redeploy by deleting pods${RESET}"
echo -e ""

# CHEAT SHEET
# the commands needed to deploy the dev operator
# --------------------
## check if there is one
## this is mostly to get some logs about what we are doing
#oc get deployment cluster-version-operator -n openshift-cluster-version
## scale it down
#oc scale deployment cluster-version-operator --replicas 0 --namespace openshift-cluster-version
## check again
#oc get deployment cluster-version-operator -n openshift-cluster-version
#
## now the operator
#oc get deployment console-operator -n openshift-console-operator
## scale it down
#oc scale deployment console-operator --replicas 0 --namespace openshift-console-operator
## check again
#oc get deployment console-operator -n openshift-console-operator
#
## check the configs
#oc get console.config.openshift.io
#oc get console.operator.openshift.io
## we want this. it should exist after merge...
#oc get clusteroperator/console
#
## should be in
## $HOME/gopaths/consoleoperator/src/github.com/openshift/console-operator
#cd $HOME/gopaths/consoleoperator/src/github.com/openshift/console-operator
#
## build the console operator binary & container image
#docker build -t quay.io/benjaminapetersen/console-operator:latest .
## push it
#docker push quay.io/benjaminapetersen/console-operator:latest
#
## deploy an alt verison of the operator with a reference to the dev version
#oc apply -f examples/05-operator-mine
