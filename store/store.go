package store

import (
	"net/http"
)

type Task struct {
	ID     string
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   string      `json:"body"`

	StatusCode   int
	LenResponse  int64
	ResultHeader http.Header
}

type Store interface {
	AddTask(t *Task) string
	GetTask(id string) (Task, bool)
	DelTask(id string) bool
	GetAllTasks() map[string]Task
}

func NewStore(storeType string) Store {
	if storeType == "map" {
		store := new(StoreMap)
		store.taskList = make(map[string]Task)
		store.idCounter = 100
		return store
	}
	return nil
}
