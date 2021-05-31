package respond

import (
	"encoding/json"
	"final_backend/models"
	"final_backend/service"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func GoogleSignIn(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var user models.User
	var response = models.Success{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	userResult, err := service.GetUser(database, user.ToQueryBson())
	// if no user found then insert a new one and return
	if userResult == nil {
		_, err := service.SaveUser(database, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		// check if insert user
		_, err = service.GetUser(database, user.ToQueryBson())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		response.Success = true
		response.Message = "Add New User"
		jsondata, _ := json.Marshal(response)
		w.Write(jsondata)
		return nil
	}
	// if already has users
	response.Success = true
	response.Message = "User Login"
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

func GetBoardById(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetBoardById queryInfo:", requestBody)

	board, err := service.GetBoardById(database, models.Board{BoardId: requestBody["boardId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	childBoards, err := service.GetChildBoardByBoardId(database, models.ChildBoard{BoardId: requestBody["boardId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response models.Board
	// var response viewModels.SentiTasksViewModel
	response.BoardId = board.BoardId
	response.BoardName = board.BoardName
	response.DomainName = board.DomainName
	response.Img = board.Img
	response.PostNum = board.PostNum

	for _, childBoard := range childBoards {
		var temp = models.ChildBoard{
			BoardId:        childBoard.BoardId,
			ChildBoardId:   childBoard.ChildBoardId,
			ChildBoardName: childBoard.ChildBoardName,
			PostNum:        childBoard.PostNum,
		}
		response.ChildBoards = append(response.ChildBoards, temp)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func GetAllBoards(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetBoardById queryInfo:", requestBody)

	boards, err := service.GetBoardsByDomain(database, models.Board{DomainName: requestBody["domainName"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response []models.Board
	for _, board := range boards {
		var temp = models.Board{
			BoardId:    board.BoardId,
			BoardName:  board.BoardName,
			DomainName: board.DomainName,
			PostNum:    board.PostNum,
			Img:        board.Img,
		}
		response = append(response, temp)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func GetAllPosts(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetBoardById queryInfo:", requestBody)

	posts, err := service.GetPostsByBoardId(database, models.Board{BoardId: requestBody["boardId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response []models.Post
	for _, post := range posts {
		var temp = models.Post{
			PostId:       post.PostId,
			BoardId:      post.BoardId,
			ChildBoardId: post.ChildBoardId,
			PostTag:      post.PostTag,
			PostTitle:    post.PostTitle,
			Author:       post.Author,
			AuthorName:   post.AuthorName,
			AuthorInfo:   post.AuthorInfo,
			Floor:        post.Floor,
			CommentNum:   post.CommentNum,
			LikeNum:      post.LikeNum,
			Time:         post.Time,
			LikedUsers:   post.LikedUsers,
		}
		response = append(response, temp)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func GetUserInfo(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetUserInfo queryInfo:", requestBody)

	user, err := service.GetUserInfoById(database, models.User{UserId: requestBody["userId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	specialties, err := service.GetSpecialtyByUserId(database, models.Specialty{UserId: requestBody["userId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	posts, err := service.GetPostsByUserId(database, models.Post{Author: requestBody["userId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	votes, err := service.GetVotesByUserId(database, models.Vote{Launcher: requestBody["userId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response models.User

	response.Name = user.Name
	response.AccessToken = user.AccessToken
	response.ImageUrl = user.ImageUrl
	response.Email = user.Email
	response.FamilyName = user.FamilyName
	response.GivenName = user.GivenName
	response.UserId = user.UserId
	response.SelfIntroduction = user.SelfIntroduction
	response.InterestGames = user.InterestGames

	for _, specialty := range specialties {
		var temp = models.Specialty{
			BoardId:   specialty.BoardId,
			BoardName: specialty.BoardName,
			UserId:    specialty.UserId,
			Score:     specialty.Score,
		}
		response.CumulateGameSpecialty = append(response.CumulateGameSpecialty, temp)
	}
	for _, post := range posts {
		var temp = models.Post{
			PostId:       post.PostId,
			BoardId:      post.BoardId,
			ChildBoardId: post.ChildBoardId,
			PostTag:      post.PostTag,
			PostTitle:    post.PostTitle,
			Author:       post.Author,
			Floor:        post.Floor,
			CommentNum:   post.CommentNum,
			LikeNum:      post.LikeNum,
			Time:         post.Time,
			LikedUsers:   post.LikedUsers,
		}
		response.PublishPost = append(response.PublishPost, temp)
	}
	for _, vote := range votes {
		var temp = models.Vote{
			VoteId:    vote.VoteId,
			Launcher:  vote.Launcher,
			BoardName: vote.BoardName,
			Img:       vote.Img,
			Agree:     vote.Agree,
			Disagree:  vote.Disagree,
			Reason:    vote.Reason,
		}
		response.LaunchNewBoard = append(response.LaunchNewBoard, temp)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func GetPostDetail(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Post
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetPostDetail queryInfo:", requestBody)

	posts, err := service.GetPostsByPostId(database, requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response []models.Post
	for _, post := range posts {
		var temp = models.Post{
			PostId:       post.PostId,
			BoardId:      post.BoardId,
			ChildBoardId: post.ChildBoardId,
			PostTag:      post.PostTag,
			PostTitle:    post.PostTitle,
			Author:       post.Author,
			AuthorName:   post.AuthorName,
			Floor:        post.Floor,
			CommentNum:   post.CommentNum,
			LikeNum:      post.LikeNum,
			Time:         post.Time,
			LikedUsers:   post.LikedUsers,
			Content:      post.Content,
			Comments:     post.Comments,
			Citations:    post.Citations,
			AuthorInfo:   post.AuthorInfo,
		}
		response = append(response, temp)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func Post(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Post

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	_, err = service.InsertPost(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func EditPost(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Post

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.UpdatePost(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func PostComment(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Comment

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	_, err = service.InsertComment(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func EditComment(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Comment

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.UpdateComment(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func EditUserInfo(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.User

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.UpdateUser(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func LikePost(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Post

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.LikePost(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func LikeComment(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Comment

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.LikeComment(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func AddSpecialty(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Specialty

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.AddSpecialty(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func MakeCitation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Citation

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	_, err = service.InsertCitation(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func GetVote(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetUserInfo")

	votes, err := service.GetVote(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	var response []models.Vote
	for _, vote := range votes {
		user, err := service.GetUserInfoById(database, models.User{UserId: vote.Launcher})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		vote.LauncherInfo = *user
		response = append(response, vote)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func GetVoteDetail(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetVoteDetail queryInfo:", requestBody)

	vote, err := service.GetVoteById(database, models.Vote{VoteId: requestBody["voteId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	jsondata, _ := json.Marshal(vote)
	_, _ = w.Write(jsondata)
	return nil
}

func Vote(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Vote API Error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("Vote queryInfo:", requestBody)

	res, err := service.Vote(database, requestBody)
	if err != nil {
		log.Println("UpdateVote Error", res)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// Get Vote to check whether the update is done
	response, err := service.GetVoteById(database, models.Vote{VoteId: requestBody["voteId"].(string)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func LaunchVote(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("LaunchVote queryInfo:", requestBody)

	voteId, err := service.LaunchVote(database, requestBody)
	if err != nil {
		log.Println("LaunchVote Error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// Get Vote to check whether LaunchVote is done
	response, err := service.GetVoteById(database, models.Vote{VoteId: voteId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func DeletePost(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Post

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.DeletePost(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func DeleteComment(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Comment

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.DeleteComment(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func DeleteCitation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Citation

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = service.DeleteCitation(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func GetNotification(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetNotification queryInfo:", requestBody)

	notification, _ := service.GetNotification(database, models.User{UserId: requestBody["userId"]})
	if notification == nil {
		log.Println("No notification")
		jsondata, _ := json.Marshal("No notification")
		_, _ = w.Write(jsondata)
	} else {
		jsondata, _ := json.Marshal(notification)
		_, _ = w.Write(jsondata)
	}
	return nil
}
