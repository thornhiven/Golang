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
	ID    	 string
	Method 	 string   `json:"method"`
	Url    	 string   `json:"url"`
	Header   string   `json:"header"`
	Body     string   `json:"body"`
}

type TaskResult struct{
	ID 			string
	StatusCode  int
	LenResponse int64
	Header  	http.Header
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

	switch strings.ToUpper(task.Method){
		case "GET":		resp, err = http.Get(task.Url)
		case "POST":	resp, err = http.Post(task.Url,task.Header,strings.NewReader(task.Body))
		case "HEAD":	resp, err = http.Get(task.Url)
		default:fmt.Println("Unknown method")//TODO
	}


	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	taskResp := TaskResult{
		ID: task.ID,
		StatusCode:  resp.StatusCode,
		LenResponse: resp.ContentLength,
		Header: resp.Header,
	}

	res, err := json.Marshal(taskResp)
	/*if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {

	taskList=make(map[string]Task)
	var t Task
	t= Task{ID: "1",
			Method: "get",
			Url: "google.com",
			}
	taskList[t.ID]=t
	t= Task{ID: "2",
			Method: "get",
			Url: "yandex.com",
			}
	taskList[t.ID]=t
	t= Task{ID: "345",
			Method: "post",
			Url: "asd.ytr",
			}
	taskList[t.ID]=t
	port:=":8080"

	http.HandleFunc("/task", createTask)
	http.HandleFunc("/gettasklist",getTaskList)
	http.HandleFunc("/deltask",deltask)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}