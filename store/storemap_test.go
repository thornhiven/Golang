package store

import (
	"reflect"
	"testing"
)

func TestDelTask(t *testing.T) {
	s := NewStore("map")
	id := s.AddTask(&Task{
		ID:     "100",
		Method: "get",
		Url:    "google.com",
	})
	if s.DelTask(id) == false {
		t.Error("Неверный ответ при корректном идентификатор")
	} else if _, ok := s.GetTask(id); ok {
		t.Error("Запись не удалена при корректном идентификаторе")
	}
	if s.DelTask(id) == true {
		t.Error("Неверный ответ при запросе на несуществующий идентификатор")
	}
}

func TestAddTask(t *testing.T) {
	s := NewStore("map")
	task := Task{
		Method: "get",
		Url:    "google.com",
	}
	id := s.AddTask(&task)
	task.ID = id
	resp, ok := s.GetTask(id)
	if !ok {
		t.Error("Запрос не добавлен в Store")
	} else if !reflect.DeepEqual(task, resp) {
		t.Error("Неверно записана структура")
	}
}

func TestGetAllTasks(t *testing.T) {
	var s StoreMap
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
