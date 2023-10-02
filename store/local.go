package store

import (
	"bytes"
	"fmt"
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

func (s *LocalStore) Save(fileName string, file io.Reader) error {
	if fileName == "" {
		return fmt.Errorf("file name cannot be empty")
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return err
	}

	if buf.Len() == 0 {
		return fmt.Errorf("file content cannot be empty")
	}

	filePath := filepath.Join(s.DirectoryName, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *LocalStore) Get(fileName string) (io.ReadCloser, error) {
	if fileName == "" {
		return nil, fmt.Errorf("file name cannot be empty")
	}

	filePath := filepath.Join(s.DirectoryName, fileName)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}
