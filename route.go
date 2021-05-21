package main

import (
	// "log"
	"net/http"

	respond "final_backend/controller"
)

type RouteMux struct {
}

func (p *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")
	switch path {
	case "/hello":
		respond.SayhelloName(w, r)
		return
	case "/testDB":
		respond.GetTest(Database, w, r)
		return
	case "/GetBoardById":
		respond.GetBoardById(Database, w, r)
		return
	default:
		break
	}
	http.NotFound(w, r)
	return
}
