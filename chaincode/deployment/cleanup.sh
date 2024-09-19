#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

printf "\n%s# Cleaning up Fabric network: *********************************\n\n%s" "$bold" "$normal"

PWD=$(pwd)
cd ../../hlf-network || exit 1
./minifab cleanup
cd ..
rm -rf hlf-network
cd "$PWD" || exit 1

printf "\n%s# Cleanup completed: ******************************************\n\n%s" "$bold" "$normal"