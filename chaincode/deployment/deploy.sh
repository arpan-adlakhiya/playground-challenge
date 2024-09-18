#!/bin/bash

mkdir -p ../../hlf-network/vars/chaincode/playground-hlf/go

cp spec.yaml ../../hlf-network/
cp minifab ../../hlf-network/

cp -R ../* ../../hlf-network/vars/chaincode/playground-hlf/go

cd ../../hlf-network/ || exit 1

./minifab up -o bank -n playground-hlf -i 2.3 -d false -l go -v 1.0 -r true -s couchdb -e 7000