package controller

import (
	"TodoList/model"
	"encoding/json"
	"html/template"
	//	"io"
	"net/http"
	"path"
	"strconv"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = getController(w, r)
	case "POST":
		err = postController(w, r)
	case "PUT":
		err = putController(w, r)
	case "DELETE":
		err = delController(w, r)
	}
	if err != nil {
		panic(err)
	}
}
func getController(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	t, _ := template.ParseFiles("View/index.html")
	if err != nil {
		t.Execute(w, nil)
	} else {
		todo, err := model.GetTodo(id)
		if err != nil {
			return err
		} else {
			data, err := json.Marshal(&todo)
			if err != nil {
				return err
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
	return nil
}

type JSON struct {
	Msg string `json:"msg"`
}

func postController(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var todo model.Todo
	json.Unmarshal(body, &todo)
	err := todo.Create()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	data := &JSON{Msg: "Create Success"}
	msg, _ := json.Marshal(&data)
	w.Write(msg)
	return nil
}
func putController(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	todo, err := model.GetTodo(id)
	if err != nil {
		return err
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &todo)
	err = todo.Update()
	if err != nil {
		return err
	}
	return nil
}
func delController(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	todo, err := model.GetTodo(id)
	if err != nil {
		return err
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &todo)
	err = todo.Delete()
	if err != nil {
		return err
	}
	return nil
}
