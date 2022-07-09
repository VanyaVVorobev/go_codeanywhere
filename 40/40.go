package task40

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var tasks []Task

type Task struct {
	IsAsync bool     `json:"is_async"`
	Time    TaskTime `json:"time"`
}

type TaskTime struct {
	Time string `json:"value"`
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

func time(w http.ResponseWriter, r *http.Request) {
	var time []TaskTime
	for _, item := range tasks {
		time = append(time, TaskTime{Time: item.Time.Time})
	}
	json.NewEncoder(w).Encode(time)
}

func Task40() {
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "0h0m1s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "0h0m2s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "0h2m1s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "1h0m1s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "0h56m1s"}})
	tasks = append(tasks, Task{IsAsync: true, Time: TaskTime{Time: "1h0m1s"}})

	r := mux.NewRouter()
	r.HandleFunc("/task/add", add).Methods("POST")
	r.HandleFunc("/task/schedule", schedule).Methods("GET")
	r.HandleFunc("/task/time", time).Methods("GET")

	fmt.Printf("Server is started at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
