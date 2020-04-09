package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDelTask(t *testing.T) {
	s := NewStore("map")
	id, _ := s.AddTask(&Task{
		ID:     "100",
		Method: "get",
		Url:    "google.com",
	})
	require.NoError(t, s.DelTask(id), "Неверный ответ при корректном идентификаторe")
	_, ok := s.GetTask(id)
	require.Error(t, ok, "Запись не удалена при корректном идентификаторе")
	require.Error(t, s.DelTask(id), "Неверный ответ при запросе на несуществующий идентификатор")
}

func TestAddTask(t *testing.T) {
	s := NewStore("map")
	task := Task{
		Method: "get",
		Url:    "google.com",
	}
	id, _ := s.AddTask(&task)
	task.ID = id
	resp, ok := s.GetTask(id)
	require.NoError(t, ok, "Запрос не добавлен в Store")
	require.Equal(t, resp, task, "Неверно записана структура")
}

func TestGetAllTasks(t *testing.T) {
	var s StoreMap
	s.Init()
	arr := make([]Task, 0)
	task := Task{
		Method: "get",
		Url:    "google.com",
	}
	s.AddTask(&task)
	arr = append(arr, task)
	task = Task{
		Method: "post",
		Url:    "yandex.ru",
	}
	s.AddTask(&task)
	arr = append(arr, task)
	v, _ := s.GetAllTasks()
	require.Equal(t, arr, v, "Возвращаемая структура не соответсвует записанной")
}
