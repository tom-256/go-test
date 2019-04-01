package main

import (
	"testing"
)

func TestSearch(t *testing.T)  {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want :=  "this is just a test"

		assertStrings(t, got, want)
	})
	
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search( "hoge")
		want :=  "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, def)
	})

}

func assertStrings(t *testing.T, got, want string)  {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	def := "this is just a test"
	err := dictionary.Add(word, def)

	assertError(t, err, nil)
	assertDefinition(t, dictionary, word, def)
}

func assertDefinition(t *testing.T, dictionary Dictionary,word, definition string)  {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func assertError(t *testing.T, err, expected error)  {
	t.Helper()

    if err != expected {
		t.Errorf("got '%s' want '%s'", err,expected)
    }
}
