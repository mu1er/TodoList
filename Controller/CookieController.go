package controller

import (
	"net/http"
)

func ReadCookieServer(r *http.Request) (string, error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return nil, err
	} else {
		return cookie.Value, nil
	}
}

func WriteCookieServer(w http.ResponseWriter, value string) {
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    value,
		MaxAge:   86400,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func DeleteCookieServer(w http.ResponseWriter) error {
	cookie := http.Cookie{
		Name:     "_cookie",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
