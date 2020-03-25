package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"net"
)

type Task struct {
	ID     string
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   string      `json:"body"`
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
func deltask(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id, ok := q["id"]
	if !ok {
		http.Error(w, "id required", 400)
		return
	} else if _, ok := taskList[id[0]]; ok {
		delete(taskList, id[0])
		fmt.Fprintf(w, "Task deleted")
	} else {
		fmt.Fprintf(w, "Id not exist")
	}
}

//возвращает список имеющихся запросов
func getTaskList(w http.ResponseWriter, r *http.Request) {
	for _, v := range taskList {
		res, err := json.Marshal(v)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

//считывает просьбу, выполняет и заносит в список
func createTask(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	task.ID= strconv.Itoa(idCounter)
	idCounter++
	taskList[task.ID] = task

	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(task.Method), task.Url, nil)
	req.Header = task.Header
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	taskResp := TaskResult{
		ID:          task.ID,
		StatusCode:  resp.StatusCode,
		LenResponse: resp.ContentLength,
		Header:      resp.Header,
	}

	res, err := json.Marshal(taskResp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {

	var port = flag.String("port", "8080", "Port value")
	flag.Parse()
	taskList = make(map[string]Task)

	http.HandleFunc("/task", createTask)
	http.HandleFunc("/gettasklist", getTaskList)
	http.HandleFunc("/deltask", deltask)

	err := http.ListenAndServe(net.JoinHostPort("localhost", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
