#!/usr/bin/env bash

# this just deletes everything under /manifests,
# but tries to space them out a bit to avoid errors.
# in the end, it creates a custom resource to kick
# the operator into action

#FILE=examples/cr.yaml
#echo "creating ${FILE}"
#oc delete -f $FILE

# heh, space matters after comma, or bash thinks its a string :)
RUNLEVELS=( 1 2 3 4 5 6 7 8 9)

for LEVEL in "${RUNLEVELS[@]}"
do
   for FILE in `find ./manifests -name "0${LEVEL}-*"`
   do
      echo "deleting ${FILE}"
      oc delete -f $FILE
   done
done

#
#
#for FILE in `find ./manifests -name '05-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#for FILE in `find ./manifests -name '04-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#for FILE in `find ./manifests -name '03-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#for FILE in `find ./manifests -name '02-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#for FILE in `find ./manifests -name '01-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#for FILE in `find ./manifests -name '00-*'`
#do
#  echo "deleting ${FILE}"
#  oc delete -f $FILE
#done
#
#
#
