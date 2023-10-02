package store

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
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
			if err != nil {
				t.Fatal(err)
			}

			f, err := ls.Get(tc.fileName)
			if err != nil {
				t.Fatal(err)
			}

			readFile, err := io.ReadAll(f)
			if err != nil {
				t.Fatal(err)
			}

			if string(readFile) != string(tc.file) {
				t.Fatalf("expected %s, got %s", string(tc.file), string(readFile))
			}
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
			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			f, err := ls.Get(tc.fileName)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if f != nil {
				defer f.Close()
				t.Fatalf("expected nil file, got %v", f)
			}
		})
	}
}
