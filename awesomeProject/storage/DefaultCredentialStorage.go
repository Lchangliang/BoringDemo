package storage

type DefaultCredentialStorage struct {
	db map[string]string
}

func New() *DefaultCredentialStorage {
	storage := DefaultCredentialStorage{db: make(map[string]string)}
	storage.db["liuchangliang"] = "0621"
	return &storage
}

func (storage *DefaultCredentialStorage) GetPasswordById(appId string) string {
	 return storage.db[appId]
}



