/*
This file contains constants and common utility functions used throughout the project
*/
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Get a state from ledger for the key passed as argument. Return error if state does not exist
func GetState(ctx contractapi.TransactionContextInterface, key string) ([]byte, error) {
	state, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("error while getting record from world state. Id:%s ,Error:%w", key, err)
	}
	if state == nil {
		return nil, fmt.Errorf("state %w for Id:%s in world state", ErrNotExist, key)
	}

	return state, nil
}

// Get a state from ledger for the key passed as argument. Return error if state does not exist
func GetPrivateState(ctx contractapi.TransactionContextInterface, collection string, key string) ([]byte, error) {
	state, err := ctx.GetStub().GetPrivateData(collection, key)
	if err != nil {
		return nil, fmt.Errorf("error while getting record from collection %s. Id:%s ,Error:%w",
			collection, key, err)
	}
	if state == nil {
		return nil, fmt.Errorf("state %w for Id:%s in collection %s", ErrNotExist, key, collection)
	}

	return state, nil
}

func GetPrivateStateForTokenCount(
	ctx contractapi.TransactionContextInterface,
	collection string,
	key string,
) ([]byte, error) {

	state, err := ctx.GetStub().GetPrivateData(collection, key)

	if err != nil {

		return nil, fmt.Errorf("error while getting record from collection %s. Id:%s ,Error:%w",
			collection, key, err)
	}

	return state, nil
}

// Verify state does not exists on the ledger. Returns error if the state exists
func VerifyStateDoesNotExist(ctx contractapi.TransactionContextInterface, key string) error {
	state, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("error while getting record from world state. Id:%s ,Error:%w", key, err)
	}
	if state != nil {
		return fmt.Errorf("state %w for key: %s", ErrAlreadyExist, key)
	}

	return nil
}

// Verify state does not exists on the ledger. Returns error if the state exists
func VerifyPrivateStateDoesNotExist(ctx contractapi.TransactionContextInterface, collection string, key string) error {
	state, err := ctx.GetStub().GetPrivateData(collection, key)
	if err != nil {
		return fmt.Errorf("error while getting record from collection %s. Id:%s ,Error:%w",
			collection, key, err)
	}
	if state != nil {
		return fmt.Errorf("private state %w for key: %s", ErrAlreadyExist, key)
	}

	return nil
}

// Put a state on ledger. Returns the serialized JSON []byte if successful
func PutState(ctx contractapi.TransactionContextInterface, key string, v interface{}) error {
	stateStr, err := json.Marshal(v)

	if err != nil {
		return fmt.Errorf("error while marshalling state. key: %s, Error:%w", key, err)
	}

	err = ctx.GetStub().PutState(key, stateStr)
	if err != nil {

		return fmt.Errorf("error while putting state in World state. key: %s, Error: %w", key, err)
	}

	return nil
}

func PutPrivateState(ctx contractapi.TransactionContextInterface, collection string, key string, v interface{}) error {
	stateStr, err := json.Marshal(v)

	if err != nil {
		return fmt.Errorf("error while marshalling state. key: %s, Error:%w", key, err)
	}

	err = ctx.GetStub().PutPrivateData(collection, key, stateStr)
	if err != nil {
		return fmt.Errorf("error while putting state in collection %s, key: %s, Error: %w", collection, key, err)
	}

	return nil
}

func SetEvent(ctx contractapi.TransactionContextInterface, eventName string, event interface{}) error {
	eventsStr, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error while marshalling events type: %s, Error:%w", eventName, err)
	}

	err = ctx.GetStub().SetEvent(eventName, eventsStr)
	if err != nil {
		return fmt.Errorf("error while setting %s event. Error: %w", eventName, err)
	}

	return nil
}

func CreateCompositeKey(ctx contractapi.TransactionContextInterface, objectType string, id string) (string, error) {
	key, err := ctx.GetStub().CreateCompositeKey(objectType, []string{id})
	if err != nil {
		return key, fmt.Errorf("error while creating composite key for %s %s, error:%w", objectType, id, err)
	}

	return key, nil
}

func VerifyPrivateDataHash(
	ctx contractapi.TransactionContextInterface,
	collectionName string,
	tokenID string,
	receivedData interface{},
) error {
	// check hash for private data exists
	hashFromLedger, err := ctx.GetStub().GetPrivateDataHash(collectionName, tokenID)
	if err != nil {
		return fmt.Errorf("error while getting private data hash for Id: %s. from collection:%s Error:%w",
			tokenID, collectionName, err)
	}
	// verify private data hash
	receivedDataStr, err := json.Marshal(receivedData)

	if err != nil {
		return fmt.Errorf("error while marshalling received Data, %w", err)
	}

	receivedDataHash := sha256.Sum256(receivedDataStr)

	receivedDataHashBytes := receivedDataHash[:]
	if !bytes.Equal(hashFromLedger, receivedDataHashBytes) {
		return fmt.Errorf("data received does %w hash on ledger. Id %s \n ledgerHash: %v \n receivedDataHash: %v",
			ErrNotMatched, tokenID, hashFromLedger, receivedDataHashBytes)
	}
	log.Println("After receivedDataHashBytes")

	return nil
}
func VerifyPrivateDataHashWithMutex(
	ctx contractapi.TransactionContextInterface,
	collectionName string,
	tokenID string,
	receivedData interface{},
	mutex *sync.Mutex,
) error {
	// check hash for private data exists
	mutex.Lock()
	hashFromLedger, err := ctx.GetStub().GetPrivateDataHash(collectionName, tokenID)
	mutex.Unlock()
	if err != nil {
		return fmt.Errorf("error while getting private data hash for Id: %s. from collection:%s Error:%w",
			tokenID, collectionName, err)
	}
	// verify private data hash
	receivedDataStr, err := json.Marshal(receivedData)
	if err != nil {
		return fmt.Errorf("error while marshalling received Data, %w", err)
	}

	receivedDataHash := sha256.Sum256(receivedDataStr)

	receivedDataHashBytes := receivedDataHash[:]
	if !bytes.Equal(hashFromLedger, receivedDataHashBytes) {
		return fmt.Errorf("data received %w with hash on ledger. Id %s \n ledgerHash: %v \n receivedDataHash: %v",
			ErrNotMatched, tokenID, hashFromLedger, receivedDataHashBytes)
	}
	log.Println("After receivedDataHashBytes")

	return nil
}

func GetTxnTimestampString(ctx contractapi.TransactionContextInterface) string {
	timeStampppb, _ := ctx.GetStub().GetTxTimestamp()
	timeStamp := time.Unix(timeStampppb.GetSeconds(), int64(timeStampppb.GetNanos()))
	loc, _ := time.LoadLocation("")
	istTime := timeStamp.In(loc)

	return istTime.Format("2006-01-02T15:04:05.000Z")
}

func GetTransientData(ctx contractapi.TransactionContextInterface, key string) ([]byte, error) {
	transientData, err := ctx.GetStub().GetTransient()
	if err != nil {
		return nil, fmt.Errorf("error while getting transient data. %w", err)
	}

	dataStr, ok := transientData[key]
	if !ok {
		return nil, fmt.Errorf("'%s' %w in the transient map input", key, ErrNotExist)
	}

	return dataStr, nil
}

func CalTokenDataHash(
	receivedData interface{},
) ([]byte, error) {

	// verify private data hash
	receivedDataStr, err := json.Marshal(receivedData)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling received Data, %w", err)
	}

	receivedDataHash := sha256.Sum256(receivedDataStr)

	return receivedDataHash[:], nil
}
