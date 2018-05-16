package main

import (
	"TodoList/controller"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./View/static/"))))
	http.HandleFunc("/todo/", controller.TodoHandler)
	http.ListenAndServe(":8000", nil)
}
