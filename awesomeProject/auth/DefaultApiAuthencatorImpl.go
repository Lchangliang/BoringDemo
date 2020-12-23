package auth

import (
	"awesomeProject/apiRequest"
	"awesomeProject/storage"
	"awesomeProject/token"
	"errors"
)
type DefaultApiAuthencatorImpl struct {
	credentialStorage storage.CredentialStorage
}

func NewbyStorage(credentialStorage storage.CredentialStorage) *DefaultApiAuthencatorImpl {
	return &DefaultApiAuthencatorImpl{
		credentialStorage: credentialStorage,
	}
}

func New() *DefaultApiAuthencatorImpl {
	return &DefaultApiAuthencatorImpl{
		credentialStorage: storage.New(),
	}
}

func (authencator *DefaultApiAuthencatorImpl) AuthByUrl(url string) error {
	request := apiRequest.NewFromFullUrl(url)
	return authencator.Auth(request)
}

func (authencator *DefaultApiAuthencatorImpl) Auth(request *apiRequest.ApiRequest) error {
	appId := request.GetAppId()
	token_ := request.GetToekn()
	timestamp := request.GetTimestamp()
	baseUrl := request.GetBaseUrl()

	clientAuthToken := token.NewAutoToeknWithoutInterval(token_, timestamp)
	if clientAuthToken.IsExpired() {
		return errors.New("Token is expired.")
	}
	password := authencator.credentialStorage.GetPasswordById(appId)
	params := make(map[string]string)
	params["password"] = password
	params["appId"] = appId
	serverAuthToken := token.New(baseUrl, timestamp, params)
	if !serverAuthToken.Match(clientAuthToken) {
		return errors.New("Token verfication failed.")
	}
	return nil
}
