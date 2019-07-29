#!/bin/bash

echo "creating logo configmaps in openshift-config"
oc create configmap my-logo-file --from-file ~/Desktop/fake-logo.png -n openshift-config
oc create configmap my-logo-file2 --from-file ~/Desktop/fake-logo.png -n openshift-config
oc create configmap my-logo-file3 --from-file ~/Desktop/fake-logo.png -n openshift-config
oc create configmap my-logo-file4 --from-file ~/Desktop/fake-logo.png -n openshift-config
oc create configmap my-logo-file5 --from-file ~/Desktop/fake-logo.png -n openshift-config
oc create configmap my-logo-file6 --from-file ~/Desktop/fake-logo.png -n openshift-config

