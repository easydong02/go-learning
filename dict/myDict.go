package dict

import "errors"

type Dictionary map[string]string

type Money int

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d *Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		(*d)[word] = def
		return nil
	} else if err == nil {
		return errWordExists
	}

	return nil
}
