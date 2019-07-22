#!/bin/bash
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

os::log::info "Running e2e tests"

PACKAGES_TO_TEST=(
    "github.com/openshift/console-operator/test/e2e/main"
    "github.com/openshift/console-operator/test/e2e/management_state"
    "github.com/openshift/console-operator/test/e2e/logging"
    "github.com/openshift/console-operator/test/e2e/providers"
    "github.com/openshift/console-operator/test/e2e/branding"
)

PACKAGES_TO_TEST_OTHER=(
    "./test/e2e/main"
    "./test/e2e/management_state"
    "./test/e2e/logging"
    "./test/e2e/providers"
    "./test/e2e/branding"
)

for PACKAGE in "${PACKAGES_TO_TEST_OTHER[@]}"
do
    os::log::info "Testing ${PACKAGE}"
    KUBERNETES_CONFIG=${KUBECONFIG} go test -timeout 30m -v $PACKAGE
done




