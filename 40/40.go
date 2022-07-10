package task40

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_code_anywhere/utils"
	"log"
	"net/http"
	"time"
)

var tasks []Task

type Task struct {
	IsAsync bool     `json:"is_async"`
	Time    TaskTime `json:"time"`
}

type TaskTime struct {
	Value string `json:"value"`
}

func add(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err.Error())
	}
	tasks = append(tasks, t)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func Time(w http.ResponseWriter, r *http.Request) {
	var curTime []TaskTime
	for _, item := range tasks {
		curTime = append(curTime, TaskTime{Value: item.Time.Value})
	}
	json.NewEncoder(w).Encode(curTime)
}

func Task40() {
	go makeTasks()
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m8s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m9s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m10s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m11s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m12s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Value: "0h0m13s"}})

	r := mux.NewRouter()
	r.HandleFunc("/task/add", add).Methods("POST")
	r.HandleFunc("/task/schedule", schedule).Methods("GET")
	r.HandleFunc("/task/time", Time).Methods("GET")

	fmt.Printf("Server is started at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func makeTasks() {
	for {
		if len(tasks) != 0 {
			for {
				curTime := utils.ParseTimeToInt(tasks[0].Time.Value)
				time.Sleep(time.Second)
				tasks[0].Time.Value = utils.ParseIntToTime(curTime - 1)
				if utils.ParseTimeToInt(tasks[0].Time.Value) == 0 {
					tasks = append(tasks[:0], tasks[1:]...)
					break
				}
			}
		}
	}
}
