package main

import (
	"net/http"

	respond "final_backend/controller"
)

type RouteMux struct {
}

func (p *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")
	switch path {
	case "/GoogleSignIn":
		respond.GoogleSignIn(Database, w, r)
		return
	case "/EditUserInfo":
		respond.EditUserInfo(Database, w, r)
		return
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
	case "/Post":
		respond.Post(Database, w, r)
		return
	case "/LikePost":
		respond.LikePost(Database, w, r)
		return
	case "/EditPost":
		respond.EditPost(Database, w, r)
		return
	case "/PostComment":
		respond.PostComment(Database, w, r)
		return
	case "/LikeComment":
		respond.LikeComment(Database, w, r)
		return
	case "/EditComment":
		respond.EditComment(Database, w, r)
		return
	case "/GetPostDetail":
		respond.GetPostDetail(Database, w, r)
		return
	case "/MakeCitation":
		respond.MakeCitation(Database, w, r)
		return
	case "/GetVote":
		respond.GetVote(Database, w, r)
		return
	case "/GetVoteDetail":
		respond.GetVoteDetail(Database, w, r)
		return
	case "/Vote":
		respond.Vote(Database, w, r)
		return
	case "/LaunchVote":
		respond.LaunchVote(Database, w, r)
		return
	default:
		break
	}
	http.NotFound(w, r)
}
