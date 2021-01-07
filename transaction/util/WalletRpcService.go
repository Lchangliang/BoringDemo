package util

type WalletRpcService interface {
	MoveMoney(id string, fromUserId int64, toUserId int64, amount float64) string
}

type MockWalletRpcServiceOne struct {}

type MockWalletRpcServiceTwo struct {}

func (mock *MockWalletRpcServiceOne) MoveMoney(id string, fromUserId int64, toUserId int64, amount float64) string {
	return "123abc"
}

func (mock *MockWalletRpcServiceTwo) MoveMoney(id string, fromUserId int64, toUserId int64, amount float64) string {
	return ""
}
