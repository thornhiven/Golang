package store

import (
	"reflect"
	"testing"
)

func TestDelTask(t *testing.T) {
	var s TaskList
	s.Init()
	s.taskList["100"] = Task{
		ID:     "100",
		Method: "get",
		Url:    "google.com",
	}
	if s.DelTask("100") == false {
		t.Error("Неверный ответ при корректном идентификатор")
	} else if _, ok := s.taskList["100"]; ok {
		t.Error("Запись не удалена при корректном идентификаторе")
	}
	if s.DelTask("100") == true {
		t.Error("Неверный ответ при запросе на несуществующий идентификатор")
	}
}

func TestAddTask(t *testing.T) {
	var s TaskList
	s.Init()
	task := Task{
		Method: "get",
		Url:    "google.com",
	}
	id := s.AddTask(&task)
	task.ID = id
	if !reflect.DeepEqual(task, s.taskList[id]) {
		t.Error("Неверно записана структура")
	}
}

func TestGetAllTasks(t *testing.T) {
	var s TaskList
	s.Init()
	task := Task{
		Method: "get",
		Url:    "google.com",
	}
	s.AddTask(&task)
	task = Task{
		Method: "post",
		Url:    "yandex.ru",
	}
	s.AddTask(&task)

	if !reflect.DeepEqual(s.taskList, s.GetAllTasks()) {
		t.Error("Возвращаемая структура не соответсвует записанной")
	}
}
