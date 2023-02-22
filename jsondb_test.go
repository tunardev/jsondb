package jsondb

import (
	"os"
	"testing"
)

const jsonPath = "test.json"

func TestSet(t *testing.T) {
	defer os.Remove(jsonPath)
	db, err := New(jsonPath)
	if err != nil {
		t.Errorf("New() failed: %v", err)
	}

	// Test Set() method
	if err := db.Set("key1", "value1"); err != nil {
		t.Errorf("Set() failed: %v", err)
	}
	if err := db.Set("key2.subkey", 123); err != nil {
		t.Errorf("Set() failed: %v", err)
	}

	// Test Get() method
	val, err := db.Get("key1")
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if val != "value1" {
		t.Errorf("Get() returned %v, want %v", val, "value1")
	}

	val, err = db.Get("key2.subkey")
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if val != 123 {
		t.Errorf("Get() returned %v, want %v", val, 123)
	}
}

func TestAdd(t *testing.T) {
	defer os.Remove(jsonPath)

	db, err := New(jsonPath)
	if err != nil {
		t.Errorf("New() failed: %v", err)
	}

	// Test Add() method
	db.Add("key4", 1)
	val, _ := db.Get("key4")
	if val != 1 {
		t.Errorf("Add() failed: %v", val)
	}

	db.Add("key4", 2)
	val, _ = db.Get("key4")
	if val != 3 {
		t.Errorf("Add() failed: %v", val)
	}

	// Test Sub() method
	db.Sub("key4", 2)
	val, _ = db.Get("key4")
	if val != 1 {
		t.Errorf("Sub() failed: %v", val)
	}

	db.Sub("key4", 1)
	val, _ = db.Get("key4")
	if val != 0 {
		t.Errorf("Sub() failed: %v", val)
	}
}

func TestPush(t *testing.T) {
	defer os.Remove(jsonPath)

	db, err := New(jsonPath)
	if err != nil {
		t.Errorf("New() failed: %v", err)
	}

	// Test Push() method
	if err := db.Push("key3", "value1"); err != nil {
		t.Errorf("Push() failed: %v", err)
	}
	if err := db.Push("key3", "value2"); err != nil {
		t.Errorf("Push() failed: %v", err)
	}

	if !db.Has("key3") {
		t.Errorf("Has() failed: key3 not found")
	}

	if err := db.Delete("key3"); err != nil {
		t.Errorf("Delete() failed: %v", err)
	}
	if db.Has("key3") {
		t.Errorf("Delete() failed: key3 still exists")
	}
}

