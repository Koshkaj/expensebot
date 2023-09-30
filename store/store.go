package store

import (
	"errors"
	"io"

	"github.com/koshkaj/expensebot/config"
)

// We could have multiple storers, gcloud, azure, aws, nfs, local, etc.
type Storer interface {
	Save(string, io.Reader) error
	Get(string) (io.ReadCloser, error)
}

func CreateFileStore(storeType string) (Storer, error) {
	switch storeType {
	case "local":
		cfg := &config.StoreConfig{
			DirectoryName: "fileStorage",
		}
		return NewLocalStore(cfg), nil
	default:
		return nil, errors.New("invalid store type")
	}

}
