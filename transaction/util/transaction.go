package util

import (
	"time"
	"strings"
	"errors"
)

type STATUS uint32

const (
	TO_BE_EXECUTED STATUS = 0
	EXECUTED STATUS = 1
	EXPIRED STATUS = 2
	FAILED STATUS = 4
)

type Transaction struct {
	Id string
	BuyerId int64
	SellerId int64
	ProductId int64
	OrderId int64
	CreateTimestamp int64
	Amount float64
	Status STATUS
	WalletTransactionId string
	walletRpcService WalletRpcService
	lock TransactionLock
}

func (transaction *Transaction) setWalletRpcService(walletRpcService WalletRpcService) {
	transaction.walletRpcService = walletRpcService
}

func (transaction *Transaction) setTransactionLock(lock TransactionLock) {
	transaction.lock = lock
}


func New(preAssignedId string, buyerId int64, sellerId int64, productId int64, orderId int64) *Transaction {
	transaction := &Transaction{
		BuyerId: buyerId,
		SellerId: sellerId,
		ProductId: productId,
		OrderId: orderId,
		Status: TO_BE_EXECUTED,
		CreateTimestamp: time.Now().UnixNano() / 1e6,
	}
	idGenerator := GetIdGenerator()
	if len(preAssignedId) != 0 {
		transaction.Id = preAssignedId
	} else {
		transaction.Id = idGenerator.GenerateTransactionId()
	}
	if !strings.HasPrefix(transaction.Id, "t_") {
		transaction.Id = "t_" + transaction.Id
	}
	return transaction
}

func (transaction *Transaction) isExpired() bool {
	executionInvokedTimestamp := time.Now().UnixNano() / 1e6
	return executionInvokedTimestamp - transaction.CreateTimestamp > 14*24*60*60*1000
}

func (transaction *Transaction) execute() (bool, error) {
	if transaction.BuyerId == 0 || (transaction.SellerId == 0 || transaction.Amount < 0.0) {
		return false, errors.New("InvalidTransactionException")
	}
	if transaction.Status == EXECUTED {
		return true, nil
	}
	isLocked := false
	defer func() {
		if (isLocked) {
			transaction.lock.UnLock(transaction.Id)
		}
	}()
	isLocked = transaction.lock.Lock(transaction.Id)
	if !isLocked {
		return false, nil
	}
	if transaction.Status == EXECUTED {
		return true, nil
	}
	if transaction.isExpired() {
		transaction.Status = EXPIRED
		return false, nil
	}
	walletTransactionId := transaction.walletRpcService.MoveMoney(transaction.Id, transaction.BuyerId, transaction.SellerId, transaction.Amount)
	if len(walletTransactionId) != 0 {
		transaction.WalletTransactionId = walletTransactionId
		transaction.Status = EXECUTED
	} else {
		transaction.Status = FAILED
		return false, nil
	}
	return true, nil
}

type TestTransaction struct {
	Transaction
}

func (test *TestTransaction) isExpired() bool {
	return true
}