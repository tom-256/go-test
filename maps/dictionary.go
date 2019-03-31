package main

type Dictionary map[string]string

func (d Dictionary)Search(dic map[string]string, word string) string {
	return dic[word]
}
