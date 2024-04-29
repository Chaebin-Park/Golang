package mydict

import "errors"

// "Dictionary" is alias of type "map[string]string"
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

func (dict Dictionary) Search(word string) (string, error) {
	value, exists := dict[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (dict Dictionary) Add(word, def string) error {
	_, err := dict.Search(word)
	
	switch err {
	case errNotFound:
		dict[word] = def
	case nil:
		return errWordExists
	}
	return nil
}