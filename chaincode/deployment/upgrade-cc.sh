#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

if [[ -z "$1" ]]; then
  printf "\n%sPlease enter the chaincode version for upgrade\n\n" "${bold}"
fi

VERSION=$1

CC_NAME=$(cat chaincode.txt)
if [[ -z "$CC_NAME" ]]; then
  printf "\n%sPlease specify chaincode name in chaincode.txt file\n" "${bold}"
  exit 1
fi

printf "\n%s# Upgrading chaincode to version %s: **************************\n\n%s" "$bold" "$VERSION" "$normal"

cd ../../hlf-network/vars/chaincode/"$CC_NAME"/go/smartcontracts || exit 1
cp -R ../../../../../../chaincode/smartcontracts/* .
cd ../../../../../
./minifab ccup -v "$VERSION" -n "$CC_NAME" -l go

printf "\n%s# Chaincode upgrade successfully: ******************************\n\n%s" "$bold" "$normal"
