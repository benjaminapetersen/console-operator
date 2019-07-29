#!/usr/bin/env bash

YELLOW="\e[93m"
RESET="\e[97m"

# note that this will kill CVO long term
# this isn't really necessary.  if you just kill the namespace
# alone, CVO will eventually bring it back, but with a delay.  Its long
# enough to trigger the failing condition.
echo -e "${YELLOW}ensure cluster-verion-operator --replicas 0${RESET}"
oc scale deployment cluster-version-operator --replicas 0 -n openshift-cluster-version
echo -e "${YELLOW}delete console namespace${RESET}"
oc delete namespace openshift-console
echo "Now, check operator logs & clusteroperator status"
# oc logs -f $(oc get --no-headers pods -o custom-columns=:metadata.name
# oc get -w clusteroperator console -o yaml
