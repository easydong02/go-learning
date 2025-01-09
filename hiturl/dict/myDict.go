package dict

import "errors"

type Dictionary map[string]string

type Money int

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")
var errCantUpdate = errors.New("Can't update non-existing word")

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

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
		return nil
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
