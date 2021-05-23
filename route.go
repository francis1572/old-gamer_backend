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
	case "/GetBoardById":
		respond.GetBoardById(Database, w, r)
		return
	case "/GetAllBoards":
		respond.GetAllBoards(Database, w, r)
		return
	case "/GetAllPosts":
		respond.GetAllPosts(Database, w, r)
		return
	case "/GetUserInfo":
		respond.GetUserInfo(Database, w, r)
		return
	default:
		break
	}
	http.NotFound(w, r)
	return
}
