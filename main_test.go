package main

import (

	//`bytes`
	//"encoding/json"

	//`net/http`
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	//`github.com/stretchr/testify/assert`

	"go-task/store"
)
import "github.com/stretchr/testify/require"

var tasKList = store.NewStore("map")

/*
func TestCreateTask(t *testing.T) {
	e := echo.New()
	body := bytes.NewBufferString(`{ "ID": "100" },
										"method":"get",
										"Url":"google.com"`)

	req := httptest.NewRequest("POST", "/task", body)
	req.Header.Set("HeaderContentType", "MIMEApplicationJSON")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	resp:=createTask(c)
	if assert.NoError(t, resp) {
		require.Equal(t, http.StatusOK, rec.Code)
		//require.Equal(t, userJSON, rec.Body.String())
	}
}
*/

func TestGetAllTasks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/gettasklist", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	er := getTaskList(c)
	tl, _ := tasKList.GetAllTasks()
	require.Equal(t, tl, er)
}
