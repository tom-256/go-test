package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary)Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition,nil
}

func (d Dictionary)Add(key, value string) error {
	_, err:=d.Search(value)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

