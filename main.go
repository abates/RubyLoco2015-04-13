package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type todo struct {
	Id   uint16
	Task string
}

var id uint16
var todoList map[string]todo

func init() {
	todoList = make(map[string]todo)
	id = 0
}

func renderJson(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func todoCrud(w http.ResponseWriter, r *http.Request) {
	var id string

	urlTokens := strings.Split(strings.TrimPrefix(r.URL.Path, "/todo"), "/")
	if len(urlTokens) > 1 {
		id = urlTokens[1]
	} else {
		http.Error(w, "Must supply an ID", http.StatusBadRequest)
		return
	}

	if t, ok := todoList[id]; ok {
		if r.Method == "GET" {
			renderJson(w, t)
		} else if r.Method == "DELETE" {
			delete(todoList, id)
		}
		return
	}
	http.NotFound(w, r)
	return
}

func todoQueryCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t todo
		err := decoder.Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id += 1
		t.Id = id
		todoList[fmt.Sprintf("%d", id)] = t
	} else if r.Method == "GET" {
		/* convert list to an array for angular $resource */
		todoListSlice := make([]todo, 0)
		for _, t := range todoList {
			todoListSlice = append(todoListSlice, t)
		}
		renderJson(w, todoListSlice)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/todo", http.HandlerFunc(todoQueryCreate))
	http.Handle("/todo/", http.HandlerFunc(todoCrud))
	http.Handle("/", http.FileServer(http.Dir("app")))

	http.ListenAndServe(":8080", nil)
}
