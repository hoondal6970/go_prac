package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word aleady exists")
	errCantupdate = errors.New("That can't update")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add for a word to a dictionary
func (d Dictionary) Add(word, def string) error {
	if _, err := d.Search(word); err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

//Update the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantupdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
