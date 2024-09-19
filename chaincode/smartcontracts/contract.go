package Contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

type SampleStruct struct {
	Key1 string  `json:"key1"`
	Key2 int     `json:"key2"`
	Key3 float64 `json:"key3"`
}

func (s *Contract) SampleTransaction(ctx contractapi.TransactionContextInterface,
	param1 string,
	param2 int,
	param3 float64,
) error {
	sampleDoc := SampleStruct{
		Key1: param1,
		Key2: param2,
		Key3: param3,
	}

	docAsBytes, _ := json.Marshal(sampleDoc)

	return ctx.GetStub().PutState(param1, docAsBytes) // Using param1 as the key to store the document
}

func (s *Contract) SampleQuery(ctx contractapi.TransactionContextInterface,
	param1 string,
) ([]SampleStruct, error) {
	queryString := fmt.Sprintf(`{"selector":{"key1":"%s"}}`, param1) // Using param1 as the key to retrieve the stored document

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var sampleDocs []SampleStruct
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var sampleDoc SampleStruct
		err = json.Unmarshal(queryResponse.Value, &sampleDoc)
		if err != nil {
			return nil, err
		}
		sampleDocs = append(sampleDocs, sampleDoc)
	}
	return sampleDocs, nil
}
