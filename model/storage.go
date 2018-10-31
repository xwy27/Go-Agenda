// Package model implements the data models
package model

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Storage is the type which represents the
// entity to store your data
// - filePath: the path of file in which your
//             data is stored
type Storage struct {
	filePath string
}

// Load is the function of a storage that
// load the data to v from the file specificed
// by storage
func (storage *Storage) load(v interface{}) error {
	file, err := os.Open(storage.filePath)
	if err != nil {
		return errors.New("Error occuried while opening file to load:\n" + err.Error())
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(v)
	if err != nil && err != io.EOF {
		return errors.New("Error occuried while loading JSON file:\n" + err.Error())
	}
	return nil
}

// write is the function of a storage that
// write the data to the file specificed by
// storage from v
func (storage *Storage) write(v interface{}) error {
	file, err := os.Open(storage.filePath)
	if err != nil {
		return errors.New("Error occuried while opening file to write:\n" + err.Error())
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(v)
	if err != nil {
		return errors.New("Error occuried while writing JSON file:\n" + err.Error())
	}
	return nil
}
