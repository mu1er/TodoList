package main

import (
	"net/http"
	"todo/controller"
)

func main() {
	http.HandleFunc("/todo/", controller.TodoHandler)
	http.ListenAndServe(":8000", nil)
}
