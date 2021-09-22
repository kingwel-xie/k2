#!/bin/bash
kubectl create ns kobh
kubectl create configmap nginx-frontend --from-file=./default.conf -n kobh
