#!/usr/bin/env bash


# give me a second before starting
sleep 2

resources[0]="service"
resources[1]="route"
resources[2]="deployment"
# resources[3]="pod" - pods have names console-<hash>
# resources[4]="oauth" - not in namespace
# resources[5]="configmap" - we have 2, neither named "console"
# resources[5]="secret" - we have 2, neither named "console"

size=${#resources[@]}

# ok, hammer our resources & break a bunch of stuff arbitrarily
for i in {1..5};
do
    resource=$(($RANDOM % $size))

    echo -n "Deleting (${i}) ";
    echo "${resource} ${resources[resource]}"
    # date ;
    oc delete ${resources[resource]} console -n openshift-console
    sleep 0.05;
done
