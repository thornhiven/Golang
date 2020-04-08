package main

import (

	//`bytes`
	//"encoding/json"

	`bytes`
	`net/http`

	//`net/http`
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	`github.com/stretchr/testify/assert`

	//`github.com/stretchr/testify/assert`

	//"go-task/store"

	`go-task/store`
)
import "github.com/stretchr/testify/require"

func TestCreateTask(t *testing.T) {
	taskList = store.NewStore("map")
	e := echo.New()
	body := bytes.NewBufferString(`{ "ID": "100" },
										"method":"get",
										"Url":"google.com"`)

	req := httptest.NewRequest("POST", "/task", body)
	req.Header.Set("HeaderContentType", "MIMEApplicationJSON")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	resp:=createTask(c)
	tl, _:=taskList.GetAllTasks()
	if assert.NoError(t, resp) {
		require.Equal(t, http.StatusOK, rec.Code, tl)
		//require.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetAllTasks(t *testing.T) {
	taskList = store.NewStore("map")
	e := echo.New()
	req := httptest.NewRequest("GET", "/gettasklist", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	er := getTaskList(c)
	tl, _ := taskList.GetAllTasks()
	require.Equal(t, tl, er)
}
