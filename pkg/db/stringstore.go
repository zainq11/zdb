package db

type StringStore interface {
	set(k string, v string) error
	get(k string) (string, error)
	size() int
}

type mapBasedStore struct {
	store map[string]string
}

func (s mapBasedStore) set(k string, v string) error {
	s.store[k] = v
	return nil
}

func (s mapBasedStore) get(k string) (string, error) {
	return s.store[k], nil
}

func (s mapBasedStore) size() int {
	return len(s.store)
}

func newStringStore() StringStore {
	return &mapBasedStore{store: make(map[string]string)}
}
