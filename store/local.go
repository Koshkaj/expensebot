package store

import (
	"io"
	"os"
	"path/filepath"

	"github.com/koshkaj/expensebot/config"
)

type LocalStore struct {
	DirectoryName string
}

func NewLocalStore(cfg *config.StoreConfig) Storer {
	return &LocalStore{
		DirectoryName: cfg.DirectoryName,
	}
}

func (ls *LocalStore) Save(fileName string, file io.Reader) error {
	path := filepath.Join(ls.DirectoryName, fileName)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}

func (ls *LocalStore) Get(fileName string) (io.ReadCloser, error) {
	path := filepath.Join(ls.DirectoryName, fileName)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
