package main

import "errors"

type Dictionary map[string]string

func (d Dictionary)Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition,nil
}

func (d Dictionary)Add(key, value string) {
	d[key] = value
}
