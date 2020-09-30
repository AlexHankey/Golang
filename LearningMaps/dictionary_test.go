package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}
func TestAdd(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
func assertDefinition(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func TestUpdate(t *testing.T)  {
	word := "test"
	definition := "this si just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Update(word, NewDefinition)
	assertDefinition(t, dictionary, word, newDefinition)
}
