package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"strings"
)

type Task struct{
	ID string
	Method string `json:"method"`
	Url string `json:"url"`
	TaskResult
}

type TaskResult struct{
	ID string
	StatusCode  int
	LenResponse int
	Headers string
}


var taskList map[string]Task
var id int=100

func deltask(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id, err := q["id"]
	if !err {
		http.Error(w, "", 400)
		return
	}else if _, ok:=taskList[id[0]] ; ok{
			delete(taskList,id[0])
			fmt.Fprintf(w,"Task deleted")
		}else{
			fmt.Fprintf(w,"Id not exist")
		}
}

//todo сделать постранично
func getTaskList(w http.ResponseWriter, r *http.Request){
	for _,v:=range taskList{
		res, err := json.Marshal(v)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)

	}
	
}

func createTask(w http.ResponseWriter, r *http.Request){
	
	body, err := ioutil.ReadAll(r.Body)
	
	defer r.Body.Close()
	
	if err != nil {
		panic(err)
	}

	var task Task
	err=nil
	
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//todo сделать нормальный генератор
	

	task.ID=strconv.Itoa(id)
	id++
	strings.ToUpper(task.Method)
	taskList[task.ID] = task
	
	var resp *http.Response
	
	err=nil
	if strings.ToUpper(task.Method) == "GET" {
		resp, err = http.Get(task.Url)
	} else {

		// todo 
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	
	err=nil
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	
	taskResp := TaskResult{
		ID: task.ID,
		StatusCode:  resp.StatusCode,
		LenResponse: len(respBody),
		//Headers:     "todo"
	}
	
    err=nil
	res, err := json.Marshal(taskResp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)	
}

func main() {
	
	taskList=make(map[string]Task)
	var t Task
	t.ID="1"
	t.Method="get"
	t.Url="google.com"
	taskList[t.ID]=t
	t.ID="2"
	t.Method="post"
	t.Url="yandex.com"
	port:=":8080"
	taskList[t.ID]=t
	t.ID="345"
	t.Method="get"
	t.Url="goe.com"
	taskList[t.ID]=t
	
	http.HandleFunc("/task", createTask)
	http.HandleFunc("/gettasklist",getTaskList)
	http.HandleFunc("/deltask",deltask)

	err := http.ListenAndServe(port, nil) 
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}