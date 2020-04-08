package store

import (
	"errors"
	"strconv"
)

type StoreMap struct {
	idCounter int
	taskList  map[string]Task
}

func (list *StoreMap) Init() {
	list.taskList = make(map[string]Task)
	list.idCounter = 100
}

func (list *StoreMap) GetTask(id string) (Task, error) {
	task, ok := list.taskList[id]
	if ok {
		return task, nil
	} else {
		return task, errors.New("Данный id не существует")
	}
}

func (list *StoreMap) DelTask(id string) error {
	if _, ok := list.taskList[id]; ok {
		delete(list.taskList, id)
		return nil
	} else {
		return errors.New("Данный id не существует")
	}
}

func (list *StoreMap) AddTask(task *Task) (string, error) {
	task.ID = strconv.Itoa(list.idCounter)
	list.idCounter++
	list.taskList[task.ID] = *task
	return task.ID, nil
}

func (list *StoreMap) GetAllTasks() ([]Task, error) {
	s := make([]Task, 0)
	for _, v := range list.taskList {
		s = append(s, v)
	}
	return s, nil
}
