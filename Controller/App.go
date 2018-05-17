package controller

import (
	"net/http"
)

func Run() {
	InitializeStatic("admin")
	InitializeStatic("default")
	Routes()
	http.ListenAndServe(":8000", nil)
}

func Routes() {
	http.HandleFunc("/todo/", TodoHandler)
}
