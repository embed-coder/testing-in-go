package app

import "net/http"

type App struct {
	Message string
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(a.Message))
}
