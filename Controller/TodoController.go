package controller

import (
	"TodoList/model"
	"log"
	"net/http"
	"strconv"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	value, _ := ReadCookieServer(r)
	vals := r.URL.Query()
	page := vals.Get("page")
	p, _ := strconv.Atoi(page)
	todos, err := model.GetAllTodo()
	if err != nil {
		log.Fatalln("Get All Todos Error", err)
	}
	paginator := model.Paginator(todos, p, 1)
	if value == "" {
		Render(w, paginator, "default", "index", "home", "nav.private", "footer")
	} else {
		Render(w, paginator, "default", "index", "home", "nav", "footer")
	}
}
