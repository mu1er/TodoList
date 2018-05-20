package controller

import (
	"TodoList/model"
	"log"
	"net/http"
	"strconv"
)

func AdminTodoListdController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	vals := r.URL.Query()
	pagenow, _ := strconv.Atoi(vals.Get("page"))
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := model.GetUser(cookie)
		if err != nil {
			log.Fatalln("Get User Error", err)
		}
		todos, err := model.GetUserTodo(user.Id)
		if err != nil {
			log.Fatalln("Get Todos Error", err)
		}
		paginator := model.Paginator(todos, pagenow, 10)
		Render(w, paginator, "admin", "default", "todolist", "nav.private")
	}
}
