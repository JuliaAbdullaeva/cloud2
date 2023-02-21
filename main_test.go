package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	const key = "crete-key"
	const value = "create-value"
	defer delete(store, key)

	var alreadyContains bool
	var val interface{}
	_, alreadyContains = store[key]
	if alreadyContains {
		t.Error("key/value already exist")
	}

	err := Put(key, value)
	if err != nil {
		t.Error(err)
	}

	val, alreadyContains = store[key]
	if !alreadyContains {
		t.Error("error create")
	}
	if val != value {
		t.Error("val/value missmatch")
	}

}

func TestGet(t *testing.T) {
	const key = "read-key"
	const value = "read-value"

	var valGet interface{}
	var errGet error
	defer delete(store, key)

	store[key] = value
	valGet, errGet = Get(key)
	if errGet != nil {
		t.Error("error Get", errGet)
	}
	if valGet != value {
		t.Error("val/value mismatch")
	}
}

func TestDelete(t *testing.T) {
	const key = "delete-key"
	const value = "delete-value"
	defer delete(store, key)

	store[key] = value
	Delete(key)
	_, contains := store[key]
	if contains {
		t.Error("error delete")
	}
}
