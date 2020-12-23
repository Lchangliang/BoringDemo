package apiRequest

import (
	"strconv"
	"strings"
)

type ApiRequest struct {
	baseUrl string
	appId string
	token string
	timestamp int64
}

func NewFromFullUrl(url string) *ApiRequest {
	url_slice := strings.Split(url, "&")
	timestamp, _ := strconv.ParseInt(url_slice[3], 10, 64)
	return &ApiRequest{
		url_slice[0],
		url_slice[1],
		url_slice[2],
		timestamp,
	}
}

func (request *ApiRequest) GetBaseUrl() string {
	return request.baseUrl
}

func (request *ApiRequest) GetAppId() string {
	return request.appId
}

func (request *ApiRequest) GetToekn() string {
	return request.token
}

func (request *ApiRequest) GetTimestamp() int64 {
	return request.timestamp
}