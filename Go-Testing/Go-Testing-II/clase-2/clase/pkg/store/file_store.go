package store

import (
	"encoding/json"
	"os"
)

type FileStore struct {
	FilePath string
}

func NewStore(fileName string) *FileStore {
	return &FileStore{FilePath: fileName}
}

func (store *FileStore) Write(data interface{}) (err error) {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(
		store.FilePath,
		fileData,
		0644,
	)
}

func (store *FileStore) Read(data interface{}) (err error) {
	file, err := os.ReadFile(store.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, data)
}
