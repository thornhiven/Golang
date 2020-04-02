package store

import "strconv"

type StoreMap struct {
	idCounter int
	taskList  map[string]Task
}

func (list *StoreMap) Init() {
	list.taskList = make(map[string]Task)
	list.idCounter = 100
}

func (list *StoreMap) GetTask(id string) (Task, bool) {
	task, ok := list.taskList[id]
	return task, ok
}

func (list *StoreMap) DelTask(id string) bool {
	if _, ok := list.taskList[id]; ok {
		delete(list.taskList, id)
		return true
	} else {
		return false
	}
}

func (list *StoreMap) AddTask(task *Task) string {
	task.ID = strconv.Itoa(list.idCounter)
	list.idCounter++
	list.taskList[task.ID] = *task
	return task.ID
}

func (list *StoreMap) GetAllTasks() map[string]Task {
	return list.taskList
}
