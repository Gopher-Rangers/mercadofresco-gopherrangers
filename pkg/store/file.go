package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data any) error
	Read(data any) error
}

type Type string

const (
	FileType Type = "file"
)

type FileStore struct {
	FileName string
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}

func (fs *FileStore) Write(data any) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data any) error {
	fileData, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, &data)
}
