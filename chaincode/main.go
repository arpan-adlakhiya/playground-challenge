package main

import (
	"log"

	smartcontracts "github.com/playground-hlf/chaincode/smartcontracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(smartcontracts.Contract)

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		log.Printf("Error creating chaincode: %s", err.Error())
	}

	if err := chaincode.Start(); err != nil {
		log.Printf("Error starting chaincode: %s", err.Error())
	}
}
