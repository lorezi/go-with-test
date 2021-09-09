package main

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("golang has practically helped me to understand those theoretical concepts in cs")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}