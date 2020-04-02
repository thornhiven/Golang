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

type TaskList struct {
	idCounter int
	taskList  map[string]Task
}

/*
type Store interface{
	Init()
	AddTask()
	DelTask(id string)
	GetAllTasks()
}*/

func (list *TaskList) Init() {
	list.taskList = make(map[string]Task)
	list.idCounter = 100
}

func (list *TaskList) DelTask(id string) bool {
	if _, ok := list.taskList[id]; ok {
		delete(list.taskList, id)
		return true
	} else {
		return false
	}
}

func (list *TaskList) AddTask(task *Task) string {
	task.ID = strconv.Itoa(list.idCounter)
	list.idCounter++
	list.taskList[task.ID] = *task
	return task.ID
}

func (list *TaskList) GetAllTasks() map[string]Task {
	return list.taskList
}
