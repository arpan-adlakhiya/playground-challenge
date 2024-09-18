#!/bin/bash
# Script to invoke chaincode
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=peer0.bank:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/bank/peers/peer0.bank/tls/ca.crt
export CORE_PEER_LOCALMSPID=bank
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/bank/users/Admin@bank/msp
export ORDERER_ADDRESS=orderer0.orderer:7050
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orderer/orderers/orderer0.orderer/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA \
  --tls -C mychannel -n playground-hlf  \
  --peerAddresses peer0.bank:7051 \
  --tlsRootCertFiles /vars/keyfiles/peerOrganizations/bank/peers/peer0.bank/tls/ca.crt \
  --peerAddresses peer0.clearing:7051 \
  --tlsRootCertFiles /vars/keyfiles/peerOrganizations/clearing/peers/peer0.clearing/tls/ca.crt \
  --peerAddresses peer0.investment:7051 \
  --tlsRootCertFiles /vars/keyfiles/peerOrganizations/investment/peers/peer0.investment/tls/ca.crt \
  -c '{"Args":["QueryTrades","TRD1"]}'
