#!/usr/bin/env bash

YELLOW="\e[93m"
RESET="\e[97m"

version=$(curl -s https://openshift-release.svc.ci.openshift.org/ | grep -m 1 -oP "<a class=\"text-success\" href=\"/releasestream/4.0.0-0.ci/release/.+>\K(.+\d+)")

echo -e "${YELLOW}Using version $version${RESET}"

rm -rf /tmp/install
mkdir /tmp/install
pushd /tmp/install


oc adm release extract --tools registry.svc.ci.openshift.org/ocp/release:$version
tar -xf openshift-install-linux-$version.tar.gz

./openshift-install --dir /tmp/cluster destroy cluster

rm -rf /tmp/cluster
mkdir /tmp/cluster
export KUBECONFIG=/tmp/cluster/auth/kubeconfig

echo =========================================================
echo
# print your pullsecret for easy pasting into the installer
# cat /path/to/pullsecret.json
cat $HOME/.secrets/try.openshift.com.pull.2019.apr.16.json

echo
echo ========================================================

./openshift-install --dir /tmp/cluster create cluster

popd
