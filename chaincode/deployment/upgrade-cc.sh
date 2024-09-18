#!/bin/bash

cd ../../hlf-network/vars/chaincode/playground-hlf/go/smartcontracts || exit 1

# rm -rf *

cp -R ../../../../../../chaincode/smartcontracts/* .

cd ../../../../../

./minifab ccup -v $1 -n playground-hlf -l go

