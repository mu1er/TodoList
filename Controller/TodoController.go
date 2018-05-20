package controller

import (
	"TodoList/model"
	"log"
	"net/http"
	"strconv"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	page, _ := strconv.Atoi(vals.Get("page"))
	todos, err := model.GetAllTodo()
	if err != nil {
		log.Fatalln("Get All Todos Error", err)
	}
	paginator := model.Paginator(todos, page, 1)
	cookie, _ := ReadCookieServer(r)
	if cookie != "" {
		Render(w, paginator, "default", "index", "home", "nav.private", "footer")
	} else {
		Render(w, paginator, "default", "index", "home", "nav", "footer")
	}
}
