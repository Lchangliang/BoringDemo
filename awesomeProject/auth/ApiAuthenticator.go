package auth

import "awesomeProject/apiRequest"

type ApiAuthenticator interface {
	AuthByUrl(url string) error
	Auth(Request *apiRequest.ApiRequest) error
}

