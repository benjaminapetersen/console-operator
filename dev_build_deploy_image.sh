#!/usr/bin/env bash

YELLOW="\e[93m"
RESET="\e[97m"

echo -e "${YELLOW}state of CVO:${RESET}"
oc get deployment cluster-version-operator -n openshift-cluster-version

# prune old images to keep disk from getting too full
# build the image
# push the image to the repository
# delete the pod
# namespace is explicit in each step
echo -e "${YELLOW}switch to project openshift-console-operator${RESET}"
oc project openshift-console-operator

echo -e "${YELLOW}pruning images older than 48hrs to keep ....${RESET}"
docker image prune -af --filter "until=48h"

echo -e "${YELLOW}building new container image...${RESET}"
docker build -t quay.io/benjaminapetersen/console-operator:latest .
echo -e "${YELLOW}pushing image...${RESET}"
docker push quay.io/benjaminapetersen/console-operator:latest

echo -e "${YELLOW}deleting operator pods...${RESET}"
oc get pods --namespace openshift-console-operator
oc delete configmap console-operator-lock
oc delete pod $(oc get --no-headers pods -o custom-columns=:metadata.name --namespace openshift-console-operator) --namespace openshift-console-operator

echo -e "${YELLOW}Deploying new pods...${RESET}"
oc get pods --namespace openshift-console-operator

