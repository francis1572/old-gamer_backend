package respond

import (
	"encoding/json"
	"final_backend/models"
	"final_backend/service"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
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

// example
func SayhelloName(w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	fmt.Println("query info:", queryInfo)

	jsondata, _ := json.Marshal(queryInfo)
	_, _ = w.Write(jsondata)
	return nil
}

func GetTest(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string

	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	var userId = queryInfo["userId"]
	log.Println("API getTest : ", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	tests, err := service.GetTest(database, models.Test{UserId: queryInfo["userId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	jsondata, _ := json.Marshal(tests)
	w.Write(jsondata)
	return nil
}
