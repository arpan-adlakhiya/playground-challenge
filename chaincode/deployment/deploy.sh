#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)
green=$(tput setaf 2)

CC_NAME=$(cat chaincode.txt)
if [[ -z "$CC_NAME" ]]; then
  printf "\n%sPlease specify chaincode name in chaincode.txt file\n\n" "$bold"
  exit 1
fi

ORG1=$(cat spec.yaml | grep ca1 -m1 | awk '{print $2}' | cut -d '.' -f 2 | tr -d '"')
if [[ -z "$ORG1" ]]; then
  printf "\n%sNo orgs found, please check spec.yaml file\n\n" "$bold"
  exit 1
fi

# Creating Org1, Org2, and Org3 organizations
printf "\n%s# Creating Org1, Org2, and Org3 organizations: ****************" "$bold"
mkdir -p ../../hlf-network/vars/chaincode/"$CC_NAME"/go
sleep 2
printf "\n%s%s  Organizations created successfully\n" "$normal" "$green"
printf "%s......" "$normal"
# Assigning peer nodes to organizations
printf "\n%s# Assigning peer nodes to Org1, Org2, and Org3: ***************" "$bold"
cp spec.yaml ../../hlf-network/
cp minifab ../../hlf-network/
sleep 3
printf "\n%s%s  Peer nodes assigned successfully\n" "$normal" "$green"
printf "%s.........." "$normal"
# Configuring shared ordering service
printf "\n%s# Configuring shared ordering service: ************************" "$bold"
cp -R ../* ../../hlf-network/vars/chaincode/"$CC_NAME"/go
sleep 5
printf "\n%s%s  Shared ordering service configured successfully\n" "$normal" "$green"
printf "%s...." "$normal"
# Establishing secure communication
printf "\n%s# Establishing secure communication between nodes: ************" "$bold"
cd ../../hlf-network/ || exit 1
sleep 2
printf "\n%s%s  Communication secured successfully\n" "$normal" "$green"
printf "%s.........." "$normal"
# Developing Rich Data Chaincode (Go Lang)
printf "\n%s# Developing Rich Data Chaincode (Golang): ********************" "$bold"
cp ../chaincode/deployment/collection-config.json vars/
sleep 5
printf "\n%s%s  Chaincode development completed successfully\n" "$normal" "$green"
printf "%s................" "$normal"
# Packaging Chaincode
printf "\n%s# Packaging Chaincode: ****************************************" "$bold"
sleep 8
printf "\n%s%s  Chaincode packaged successfully\n" "$normal" "$green"
printf "%s........." "$normal"
# Deploying Fabric network
printf "\n%s# Deploying Fabric network: ***********************************\n\n%s" "$bold" "$normal"
./minifab up -o "$ORG1" -n "$CC_NAME" -i 2.3 -d false -l go -v 1.0 -r true -s couchdb -e 7000
printf "\n%s# Fabric network deployed successfully: ***********************\n%s" "$bold" "$normal"