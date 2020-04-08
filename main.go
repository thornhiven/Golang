package main

//noinspection ALL
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-task/store"
	"io/ioutil"
	"net/http"
	"strings"
)

type TaskResult struct {
	ID          string
	StatusCode  int
	LenResponse int64
	Header      http.Header
}

var taskList store.Store //список запросов

//удаление запроса с заданым индексом
func delTask(c echo.Context) error {
	id := c.Param("id")
	if taskList.DelTask(id) == nil {
		return c.String(http.StatusOK, "Task Deleted")
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
}

//возвращает список имеющихся запросов
func getTaskList(c echo.Context) error {
	tasks, err := taskList.GetAllTasks()
	if err != nil {
		return c.String(400, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

//считывает просьбу, выполняет и заносит в список
func createTask(c echo.Context) error {

	task := new(store.Task)
	if err := c.Bind(task); err != nil {
		return c.String(400, err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(task.Method), task.Url, nil)
	if err != nil {
		return c.String(400, err.Error())
	}
	req.Body = ioutil.NopCloser(strings.NewReader(task.Body))
	req.Header = task.Header
	resp, err := client.Do(req)

	if err != nil {
		return c.String(400, err.Error())
	}
	task.StatusCode = resp.StatusCode
	task.LenResponse = resp.ContentLength
	task.ResultHeader = resp.Header
	id, err := taskList.AddTask(task)
	if err != nil {
		return c.String(400, err.Error())
	} else {
		return c.JSON(http.StatusOK, TaskResult{
			ID:          id,
			StatusCode:  resp.StatusCode,
			LenResponse: resp.ContentLength,
			Header:      resp.Header,
		})
	}
}

func main() {
	taskList = store.NewStore("map")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/task", createTask)
	e.GET("/gettasklist", getTaskList)
	e.GET("/deltask/:id", delTask)
	e.Logger.Fatal(e.Start(":8080"))
}
