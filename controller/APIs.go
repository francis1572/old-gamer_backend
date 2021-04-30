package respond

import (
	// "context"
	"encoding/json"
	"final_backend/models"
	"final_backend/service"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	// "strconv"
	// "final_backend/models"
	// "final_backend/service"
	// "final_backend/viewModels"
	// uuid "github.com/nu7hatch/gouuid"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

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
