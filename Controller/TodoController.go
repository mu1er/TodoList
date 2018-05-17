package controller

import (
	//	"TodoList/model"
	"net/http"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, "default", "index", "home", "nav", "footer")
}
