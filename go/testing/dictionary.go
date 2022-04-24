package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrWordNotFound    = errors.New("could not find the word you were looking for")
	ErrWordExists      = errors.New("cannot add word because it already exists")
	ErrWordNotExisting = errors.New("word does not exist")
)

func (dictionary Dictionary) SearchEntry(key string) (string, error) {
	definition, ok := dictionary[key]

	if ok {
		return definition, nil
	}

	return "", ErrWordNotFound
}

func (dictionary Dictionary) AddEntry(key, definition string) error {
	_, err := dictionary.SearchEntry(key)

	if err != nil {
		dictionary[key] = definition
		return nil
	}

	return ErrWordExists
}

func (dictionary Dictionary) UpdateEntry(key, definition string) error {
	_, err := dictionary.SearchEntry(key)

	if err != nil {
		return ErrWordNotExisting
	}

	dictionary[key] = definition
	return nil
}

func (dictionary Dictionary) DeleteEntry(key string) {
	delete(dictionary, key)
}
