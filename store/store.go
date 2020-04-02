package store

import (
	"net/http"
	"strconv"
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

type Store struct {
	idCounter int
	taskList  map[string]Task
}

func (list *Store) Init() {
	list.taskList = make(map[string]Task)
	list.idCounter = 100
}

func (list *Store) DelTask(id string) bool {
	if _, ok := list.taskList[id]; ok {
		delete(list.taskList, id)
		return true
	} else {
		return false
	}
}

func (list *Store) AddTask(task *Task) string {
	task.ID = strconv.Itoa(list.idCounter)
	list.idCounter++
	list.taskList[task.ID] = *task
	return task.ID
}

func (list *Store) GetAllTasks() map[string]Task {
	return list.taskList
}
