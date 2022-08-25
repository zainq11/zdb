package db

type StringStore interface {
	set(k string, v string)
	get(k string) string
	size() int
}

type mapBasedStore struct {
	store map[string]string
}

func (s mapBasedStore) set(k string, v string) {
	s.store[k] = v
}

func (s mapBasedStore) get(k string) string {
	return s.store[k]
}

func (s mapBasedStore) size() int {
	return len(s.store)
}

func newStringStore() StringStore {
	return &mapBasedStore{store: make(map[string]string)}
}
