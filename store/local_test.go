package store

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalStore(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	ls := &LocalStore{
		DirectoryName: tmpDir,
	}

	testCases := []struct {
		fileName string
		file     []byte
	}{
		{
			fileName: "test1",
			file:     []byte("test1"),
		},
		{
			fileName: "test2",
			file:     []byte("test2"),
		},
		{
			fileName: "test3",
			file:     []byte("test3"),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestLocalStore-%s", tc.fileName), func(t *testing.T) {
			buf := bytes.NewBuffer(tc.file)
			err := ls.Save(tc.fileName, buf)
			assert.NoError(t, err)

			f, err := ls.Get(tc.fileName)
			assert.NoError(t, err)

			readFile, err := io.ReadAll(f)
			assert.NoError(t, err)

			assert.Equal(t, tc.file, readFile)
		})
	}
}

func TestLocalStoreError(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	ls := &LocalStore{
		DirectoryName: tmpDir,
	}

	testCases := []struct {
		fileName string
		file     []byte
	}{
		{
			fileName: "",
			file:     []byte("test1"),
		},
		{
			fileName: "test2",
			file:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestLocalStoreError-%s", tc.fileName), func(t *testing.T) {
			buf := bytes.NewBuffer(tc.file)
			err := ls.Save(tc.fileName, buf)
			assert.Error(t, err)

			f, err := ls.Get(tc.fileName)
			assert.Error(t, err)
			assert.Nil(t, f)
		})
	}
}
