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
	http.HandleFunc("/view", TodoView)
}
func AdminRoutes() {
	http.HandleFunc("/admin/", AdminController)
	http.HandleFunc("/admin/todolist/", AdminInfoController)
	http.HandleFunc("/admin/user/", AdminUserController)
	http.HandleFunc("/admin/createtodo/", AdminCreateController)
	http.HandleFunc("/admin/edituser/", AdminEditUserController)
	http.HandleFunc("/logout/", Logout)
}
