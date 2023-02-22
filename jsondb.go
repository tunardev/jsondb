package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Database interface {
	// Set() method sets a key-value pair in the database.
	Set(key string, value interface{}) error

	// Get() method returns the value of the given key.
	Get(key string) (interface{}, error)

	// Delete() method deletes the given key from the database.
	Delete(key string) error

	// Push() method adds the given value to the array of the given key.
	Push(key string, value interface{}) error

	// Has() method returns true if the given key exists in the database.
	Has(key string) bool

	// Add() method adds the given count to the value of the given key.
	Add(key string, count int) error

	// Sub() method subtracts the given count from the value of the given key.
	Sub(key string, count int) error
}

type database struct {
	file string
	json   map[string]interface{}
}

// New() method creates a new database.
func New(file string) (Database, error) {
	if file == "" {
		panic("file cannot be empty")
	}

	db := &database{
		file: file,
		json:   make(map[string]interface{}),
	}

	if _, err := os.Stat(db.file); os.IsNotExist(err) {
		_, err := os.Create(db.file)
		if err != nil {
			return nil, err
		}

		if err := db.save(); err != nil {
			return nil, err
		}
	} else {
		savedData, err := ioutil.ReadFile(db.file)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(savedData, &db.json); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// save() method saves the database to the file.
func (db *database) save() error {
	data, err := json.MarshalIndent(db.json, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(db.file, data, 0644)
	if err != nil {
		return err 
	}
	return nil
}

func (db *database) Set(key string, value interface{}) error {
	if key == "" || value == nil {
		return errors.New("key and value cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			data[key] = make(map[string]interface{})
		}
		data = data[key].(map[string]interface{})
	}
	data[keys[len(keys)-1]] = value

	return db.save()
}

func (db *database) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("key cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			return nil, errors.New("key not found")
		}
		data = data[key].(map[string]interface{})
	}

	return data[keys[len(keys)-1]], nil
}

func (db *database) Delete(key string) error {
	if key == "" {
		return 	errors.New("key cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			return errors.New("key not found")
		}
		data = data[key].(map[string]interface{})
	}

	delete(data, keys[len(keys)-1])
	return db.save()
} 

func (db *database) Push(key string, value interface{}) error {
	if key == "" || value == nil {
		return errors.New("key or value cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			data[key] = make(map[string]interface{})
		}
		data = data[key].(map[string]interface{})
	}

	if _, ok := data[keys[len(keys)-1]]; !ok {
		data[keys[len(keys)-1]] = make([]interface{}, 0)
	}

	data[keys[len(keys)-1]] = append(data[keys[len(keys)-1]].([]interface{}), value)
	return db.save()
}

func (db *database) Has(key string) bool {
	if key == "" {
		return false
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			return false
		}
		data = data[key].(map[string]interface{})
	}

	_, ok := data[keys[len(keys)-1]]
	return ok
}

func (db *database) Add(key string, count int) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			data[key] = make(map[string]interface{})
		}
		data = data[key].(map[string]interface{})
	}

	if _, ok := data[keys[len(keys)-1]]; !ok {
		data[keys[len(keys)-1]] = 0
	}

	data[keys[len(keys)-1]] = data[keys[len(keys)-1]].(int) + count
	return db.save()
}

func (db *database) Sub(key string, count int) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	keys := strings.Split(key, ".")
	data := db.json
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		if _, ok := data[key]; !ok {
			data[key] = make(map[string]interface{})
		}
		data = data[key].(map[string]interface{})
	}

	if _, ok := data[keys[len(keys)-1]]; !ok {
		data[keys[len(keys)-1]] = 0
	}

	data[keys[len(keys)-1]] = data[keys[len(keys)-1]].(int) - count
	return db.save()
}
