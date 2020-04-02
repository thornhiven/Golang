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
//var idCounter int = 100

//удаление запроса с заданым индексом
func deltask(c echo.Context) error {
	id := c.Param("id")
	if taskList.DelTask(id) {
		return c.String(http.StatusOK, "Task Deleted")
	} else {
		return c.String(http.StatusNotFound, "Id not found")
	}
}

//возвращает список имеющихся запросов
func getTaskList(c echo.Context) error {
	return c.JSON(http.StatusOK, taskList.GetAllTasks())
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
	id := taskList.AddTask(task)

	return c.JSON(http.StatusOK, TaskResult{
		ID:          id,
		StatusCode:  resp.StatusCode,
		LenResponse: resp.ContentLength,
		Header:      resp.Header,
	})
}

func main() {
	//taskList = make(map[string]Task)
	taskList.Init()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/task", createTask)
	e.GET("/gettasklist", getTaskList)
	e.GET("/deltask/:id", deltask)
	e.Logger.Fatal(e.Start(":8080"))
}
