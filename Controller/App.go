package controller

import (
	"net/http"
)

func Run() {
	InitializeStatic("admin")
	InitializeStatic("default")
	Routes()
	AdminRoutes()
	http.ListenAndServe(":8000", nil)
}

func Routes() {
	http.HandleFunc("/", TodoHandler)
	http.HandleFunc("/todo/", TodoHandler)
	http.HandleFunc("/signup/", Signup)
	http.HandleFunc("/login/", Login)
}
func AdminRoutes() {
	http.HandleFunc("/admin/", AdminController)
	http.HandleFunc("/admin/todolist/", AdminTodoListdController)
	http.HandleFunc("/logout/", Logout)
}
