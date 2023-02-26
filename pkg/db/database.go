package db

import "time"

type Database interface {
	GetName() string
	GetCreatedTime() time.Time
	DatabaseCommands
}

type DatabaseCommands interface {
	Get(k string) (string, error)
	Set(k string, v string) error
}

type Instance struct {
	Name        string
	CreatedTime time.Time
	stringStore StringStore
}

func (i Instance) GetName() string {
	return i.Name
}

func (i Instance) GetCreatedTime() time.Time {
	return i.CreatedTime
}

func (i Instance) Get(k string) (string, error) {
	return i.stringStore.get(k)
}

func (i Instance) Set(k string, v string) error {
	return i.stringStore.set(k, v)
}

func CreateDatabase(name string) Database {
	return &Instance{
		Name:        name,
		CreatedTime: time.Now(),
		stringStore: newStringStore(),
	}
}
