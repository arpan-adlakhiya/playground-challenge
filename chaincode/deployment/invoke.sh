#!/bin/bash

CC_NAME=$(cat chaincode.txt)
if [[ -z "$CC_NAME" ]]; then
  printf "\nPlease specify chaincode name in chaincode.txt file\n\n"
  exit 1
fi

# Record a transaction
./minifab invoke -o org1 -n "$CC_NAME" -p '"SampleTransaction","param1","param2","param3"'

# Query a transaction
./minifab invoke -o org1 -n "$CC_NAME" -p '"SampleQuery","param1"'
