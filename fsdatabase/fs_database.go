package fsdatabase

import (
	"encoding/gob"
	"os"
)

type fsDatabase[K any] struct {
	file *os.File
}

type FSDatabase[K any] interface {
	Read() (K, error)
	Write(database K) error
}

func New[K any](dbPath string) (FSDatabase[K], error) {
	instance := &fsDatabase[K]{}

	file, fileErr := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0666)
	if fileErr != nil {
		return instance, fileErr
	}

	instance.file = file

	return instance, nil
}

func (fs *fsDatabase[K]) Read() (K, error) {
	var instance K

	decodeErr := gob.NewDecoder(fs.file).Decode(&instance)
	if decodeErr != nil && decodeErr.Error() != "EOF" {
		return instance, decodeErr
	}

	return instance, nil
}

func (fs *fsDatabase[K]) Write(database K) error {
	if truncateErr := fs.file.Truncate(0); truncateErr != nil {
		return truncateErr
	}
	if _, seekErr := fs.file.Seek(0, 0); seekErr != nil {
		return seekErr
	}
	return gob.NewEncoder(fs.file).Encode(database)
}
