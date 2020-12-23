package storage

type CredentialStorage interface {
	GetPasswordById(appId string) string
}