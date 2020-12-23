package token

import (
	"strconv"
	"time"
)

const DEFAULT_EXPIRED_TIME_INTERVAL int64 = 1 * 60 * 1000  // ms

type AuthToken struct {
	token string
	createTime int64
	expiredTimeInterval int64
}

func createToken(baseUrl string, createTime int64, appId string, password string) string {
	return baseUrl + strconv.FormatInt(createTime, 10) + appId + password
}

func New(baseUrl string, createTime int64, params map[string]string) *AuthToken {
	expiredTimeIntervalStr, ok := params["expiredTimeInterval"]
	expiredTimeInterval, _ := strconv.ParseInt(expiredTimeIntervalStr, 10, 64)
	token := createToken(baseUrl, createTime, params["appId"], params["password"])
	if ok {
		return NewAutoToken(token, createTime, expiredTimeInterval)
	} else {
		return NewAutoToeknWithoutInterval(token, createTime)
	}
}
func NewAutoToeknWithoutInterval(token string, createTime int64) *AuthToken {
	return &AuthToken{
		token,
		createTime,
		DEFAULT_EXPIRED_TIME_INTERVAL,
	}
}

func NewAutoToken(token string, createTime int64, expiredTimeInterval int64) *AuthToken {
	return &AuthToken{
		token,
		createTime,
		expiredTimeInterval,
	}
}

func (token *AuthToken) GetToken() string {
	return token.token
}

func (token *AuthToken) IsExpired() bool {
	currentTime := time.Now().UnixNano() / 1e6   // ms
	return currentTime - token.createTime > token.expiredTimeInterval
}

func (token *AuthToken) Match(other *AuthToken) bool {
	return token.token == other.token
}
