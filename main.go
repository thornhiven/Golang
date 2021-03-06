package main

import (
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type TaskResult struct {
	ID          string
	StatusCode  int
	LenResponse int64
	Header      http.Header
}

var taskList map[string]Task //список запросов
var idCounter int = 100

//удаление запроса с заданым индексом
func deltask(c echo.Context) error {
	id := c.Param("id")
	if _, ok := taskList[id]; ok {
		delete(taskList, id)
		return c.String(http.StatusOK, "Task Deleted")
	} else {
		return c.String(http.StatusNotFound, "Id not found")
	}
}

//возвращает список имеющихся запросов
func getTaskList(c echo.Context) error {
	return c.JSON(http.StatusOK, taskList)
}

//считывает просьбу, выполняет и заносит в список
func createTask(c echo.Context) error {

	task:=new(Task)
	if err := c.Bind(task); err != nil {
		return c.String(400, err.Error())
	}

	idCounter++
	taskList[task.ID] = *task
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
	task.ID = strconv.Itoa(idCounter)

	return c.JSON(http.StatusOK, TaskResult{
		ID:          task.ID,
		StatusCode:  resp.StatusCode,
		LenResponse: resp.ContentLength,
		Header:      resp.Header,
	})
}

func main() {
	taskList = make(map[string]Task)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/task", createTask)
	e.GET("/gettasklist", getTaskList)
	e.GET("/deltask/:id", deltask)
	e.Logger.Fatal(e.Start(":8080"))
}
