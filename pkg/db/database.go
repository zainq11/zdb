package db

import "time"

type Instance struct {
	Name        string
	CreatedTime time.Time
	stringStore StringStore
}

func (i Instance) getName() string {
	return i.Name
}

func (i Instance) getCreatedTime() time.Time {
	return i.CreatedTime
}

func (i Instance) get(k string) string {
	return i.stringStore.get(k)
}

func (i Instance) set(k string, v string) {
	i.stringStore.set(k, v)
}

type Database interface {
	getName() string
	getCreatedTime() time.Time
}

type DatabaseCommands interface {
	get(k string) string
	set(k string, v string)
}

func CreateDatabase(name string) Database {
	return Instance{
		Name:        name,
		CreatedTime: time.Now(),
		stringStore: newStringStore(),
	}
}
