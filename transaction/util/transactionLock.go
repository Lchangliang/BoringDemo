package util

type TransactionLock interface {
	Lock(id string) bool
	UnLock(id string) bool
}
