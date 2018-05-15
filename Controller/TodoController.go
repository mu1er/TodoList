package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"todo/model"
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
	if err != nil {
		return err
	}
	todo, err := model.GetTodo(id)
	if err != nil {
		return err
	}
	data, err := json.Marshal(&todo)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	return nil
}

func postController(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	todo := new(model.Todo)
	json.Unmarshal(body, &todo)
	err := todo.Create()
	if err != nil {
		return err
	}
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
