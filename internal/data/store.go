package data

import "errors"

var ErrorNotFound = errors.New("not found")

type Store interface{
	// Get fetches the data stored in `key` and unmarshals it into `value`.
	Get(key string, value interface{}) error

	// Put marshals `value` and stores it in `key`.
	Put(key string, value interface{}) error

	// Delete removes `key`
	Delete(key string) error

	// List returns all the keys that start with `prefix`. The prefix is
	// stripped from each returned value. So if keys are ["aa", "ab", "cd"]
	// then List("a") would produce []string{"a", "b"}
	List(prefix string) ([]string, error)
}