#!/bin/bash
kubectl create ns k2
kubectl create configmap settings-admin --from-file=../../config/settings.yml -n k2
