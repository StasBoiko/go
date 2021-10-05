package main

import "fmt"

const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr are errors that can happen when interacting with the dictionary.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	fmt.Println(definition, ok)
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	// fmt.Println(err)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err

	}

	return nil
}

// Update changes the definition of a given word.
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}

func main () {
	// dictionary := Dictionary{}
	// word := "test5"
	// definition := "aaa"

	// word2 := "test5"
	// definition2 := "bbb"
	// dictionary.Add(word, definition)
	// dictionary.Add(word2, definition2)
	// fmt.Println(foo)
	// fmt.Println(foo2)
	// fmt.Println(dictionary)


	// word := "test"
	// dictionary := Dictionary{word: "test definition"}

	// fmt.Println(dictionary)
	// delete(dictionary, word)
	// fmt.Println(dictionary)
	// dictionary.Delete(word)
}
