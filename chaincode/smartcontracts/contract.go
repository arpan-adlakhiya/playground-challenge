package Contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

type Trade struct {
	TradeID   string  `json:"tradeId"`
	Symbol    string  `json:"symbol"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

type Payment struct {
	PaymentID string  `json:"paymentId"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

type Settlement struct {
	SettlementID string  `json:"settlementId"`
	TradeID      string  `json:"tradeId"`
	PaymentID    string  `json:"paymentId"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
	Status       string  `json:"status"`
}

func (s *Contract) RecordTrade(ctx contractapi.TransactionContextInterface, tradeID string, symbol string, quantity int, price float64, timestamp string, status string) error {
	trade := Trade{
		TradeID:   tradeID,
		Symbol:    symbol,
		Quantity:  quantity,
		Price:     price,
		Timestamp: timestamp,
		Status:    status,
	}
	tradeAsBytes, _ := json.Marshal(trade)
	return ctx.GetStub().PutState(tradeID, tradeAsBytes)
}

func (s *Contract) RecordPayment(ctx contractapi.TransactionContextInterface, paymentID string, sender string, receiver string, amount float64, timestamp string, status string) error {
	payment := Payment{
		PaymentID: paymentID,
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: timestamp,
		Status:    status,
	}
	paymentAsBytes, _ := json.Marshal(payment)
	return ctx.GetStub().PutState(paymentID, paymentAsBytes)
}

func (s *Contract) HandleSettlement(ctx contractapi.TransactionContextInterface, settlementID string, tradeID string, paymentID string, amount float64, timestamp string, status string) error {
	settlement := Settlement{
		SettlementID: settlementID,
		TradeID:      tradeID,
		PaymentID:    paymentID,
		Amount:       amount,
		Timestamp:    timestamp,
		Status:       status,
	}
	settlementAsBytes, _ := json.Marshal(settlement)
	return ctx.GetStub().PutState(settlementID, settlementAsBytes)
}

func (s *Contract) QueryTrades(ctx contractapi.TransactionContextInterface, tradeID string) ([]Trade, error) {
	queryString := fmt.Sprintf(`{"selector":{"tradeId":"%s"}}`, tradeID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var trades []Trade
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var trade Trade
		err = json.Unmarshal(queryResponse.Value, &trade)
		if err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}
	return trades, nil
}

func (s *Contract) QueryPayments(ctx contractapi.TransactionContextInterface, paymentID string) ([]Payment, error) {
	queryString := fmt.Sprintf(`{"selector":{"paymentId":"%s"}}`, paymentID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var payments []Payment
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var payment Payment
		err = json.Unmarshal(queryResponse.Value, &payment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (s *Contract) QuerySettlements(ctx contractapi.TransactionContextInterface, settlementID string) ([]Settlement, error) {
	queryString := fmt.Sprintf(`{"selector":{"settlementId":"%s"}}`, settlementID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var settlements []Settlement
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var settlement Settlement
		err = json.Unmarshal(queryResponse.Value, &settlement)
		if err != nil {
			return nil, err
		}
		settlements = append(settlements, settlement)
	}
	return settlements, nil
}
