package service

import (
	"context"
	"log"
	"strconv"
	"time"

	"final_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(db *mongo.Database, queryBson bson.M) (*models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = models.User{}
	cur := collection.FindOne(context.Background(), queryBson)
	// if no user then return nil
	if cur.Err() != nil {
		log.Println("Can't find user in DB")
		return nil, cur.Err()
	}
	// if has user then return
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode user Error", err)
		return nil, err
	}
	log.Println("Get user:", serviceResult)
	return &serviceResult, nil
}

func GetUsers(db *mongo.Database) ([]models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = []models.User{}
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Find user Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.User{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func SaveUser(db *mongo.Database, user models.User) (*mongo.InsertOneResult, error) {
	UserCollection := db.Collection("GUser")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Insert user Error", err)
		return nil, err
	}
	log.Println("Insert user Success", user)
	return res, nil
}

func GetBoardById(db *mongo.Database, query models.Board) (*models.Board, error) {
	BoardCollection := db.Collection("Board")
	var board models.Board
	result := BoardCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&board)
	// log.Println("Success", board)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &board, nil
}

func GetBoardsByDomain(db *mongo.Database, task models.Board) ([]*models.Board, error) {
	BoardCollection := db.Collection("Board")
	var Boards []*models.Board

	cur, err := BoardCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Board{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		Boards = append(Boards, &result)
	}
	return Boards, nil
}

func GetChildBoardByBoardId(db *mongo.Database, task models.ChildBoard) ([]*models.ChildBoard, error) {
	ChildBoardCollection := db.Collection("ChildBoard")
	var childBoards []*models.ChildBoard
	cur, err := ChildBoardCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.ChildBoard{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		childBoards = append(childBoards, &result)
	}
	return childBoards, nil
}

func GetPostsByBoardId(db *mongo.Database, task models.Board) ([]*models.Post, error) {
	PostCollection := db.Collection("Post")
	var posts []*models.Post
	cur, err := PostCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Post{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		posts = append(posts, &result)
	}
	return posts, nil
}

func GetUserInfoById(db *mongo.Database, query models.User) (*models.User, error) {
	GUserCollection := db.Collection("GUser")
	var user models.User
	result := GUserCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&user)
	// log.Println("Success", board)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &user, nil
}

func GetSpecialtyByUserId(db *mongo.Database, task models.Specialty) ([]*models.Specialty, error) {
	SpecialtyCollection := db.Collection("Specialty")
	var specialties []*models.Specialty
	cur, err := SpecialtyCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Specialty{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		specialties = append(specialties, &result)
	}

	return specialties, nil
}

func GetPostsByUserId(db *mongo.Database, task models.Post) ([]*models.Post, error) {
	PostCollection := db.Collection("Post")
	var posts []*models.Post
	cur, err := PostCollection.Find(context.Background(), task.ToQueryBson())
	// log.Println("GetUserInfo posts:", task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Post{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		posts = append(posts, &result)
	}
	// log.Println("GetUserInfo posts:", posts)
	return posts, nil
}

func GetVotesByUserId(db *mongo.Database, task models.Vote) ([]*models.Vote, error) {
	VoteCollection := db.Collection("Vote")
	var votes []*models.Vote
	cur, err := VoteCollection.Find(context.Background(), task.ToQueryBson())
	// log.Println("GetUserInfo Votes:", task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		votes = append(votes, &result)
	}
	// log.Println("GetUserInfo votes:", votes)
	return votes, nil
}

func GetPostsByPostId(db *mongo.Database, task models.Post) ([]*models.Post, error) {
	PostCollection := db.Collection("Post")
	var posts []*models.Post
	cur, err := PostCollection.Find(context.Background(), task.ToQueryBson())
	// log.Println("GetUserInfo posts:", task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	log.Println("有通過第一階段")
	for cur.Next(context.Background()) {
		result := models.Post{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		log.Println("有通過第二階段")
		user, _ := GetUserInfoById(db, models.User{UserId: result.Author})
		result.AuthorInfo = *user
		log.Println("有通過第三階段", result.AuthorInfo)
		blocks, _ := GetBlocksByFloor(db, models.Block{PostId: result.PostId, Floor: result.Floor})
		for _, block := range blocks {
			var temp = models.Block{
				PostId:   block.PostId,
				Floor:    block.Floor,
				BlockId:  block.BlockId,
				Subtitle: block.Subtitle,
				Content:  block.Content,
			}
			result.Content = append(result.Content, temp)
		}
		log.Println("有通過第四階段", result.Content)
		citations, err := GetCitesByFloor(db, models.Citation{PostId: result.PostId, Floor: result.Floor})
		if err == nil {
			for _, citation := range citations {
				var temp = models.Citation{
					PostId:     citation.PostId,
					Floor:      citation.Floor,
					CitationId: citation.CitationId,
					CitedFloor: citation.CitedFloor,
					BlockId:    citation.BlockId,
				}
				result.Citations = append(result.Citations, temp)
			}
		}
		log.Println("有通過第五階段", result.Citations)

		posts = append(posts, &result)
	}
	// log.Println("GetUserInfo posts:", posts)
	return posts, nil
}

func GetBlocksByFloor(db *mongo.Database, task models.Block) ([]*models.Block, error) {
	BlockCollection := db.Collection("Block")
	var blocks []*models.Block
	cur, err := BlockCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Block{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		blocks = append(blocks, &result)
	}
	return blocks, nil
}

func GetCitesByFloor(db *mongo.Database, task models.Citation) ([]*models.Citation, error) {
	CiteCollection := db.Collection("Citation")
	var cites []*models.Citation
	cur, err := CiteCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Citation{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		cites = append(cites, &result)
	}
	return cites, nil
}

func GetVoteById(db *mongo.Database, query models.Vote) (*models.Vote, error) {
	VoteCollection := db.Collection("Vote")
	var vote models.Vote
	result := VoteCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&vote)
	if err != nil {
		log.Println("Decode vote Error", err)
		return nil, err
	}
	return &vote, nil
}

func UpdateVote(db *mongo.Database, queryBson bson.M) (*mongo.UpdateResult, error) {
	VoteCollection := db.Collection("Vote")
	var vote models.Vote
	result := VoteCollection.FindOne(context.Background(), bson.M{"voteId": queryBson["voteId"]})
	err := result.Decode(&vote)
	if err != nil {
		log.Println("Decode vote Error", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if contains(vote.DisagreedUsers, queryBson["userId"].(string)) || contains(vote.AgreedUsers, queryBson["userId"].(string)) {
		log.Println("User voted")
		return nil, nil
	}

	var update bson.M
	// type of queryBson["decision"] is float
	if queryBson["decision"] == 0. {
		vote.DisagreedUsers = append(vote.DisagreedUsers, queryBson["userId"].(string))
		update = bson.M{"$set": bson.M{"disagree": vote.Disagree + 1, "disagreedUsers": vote.DisagreedUsers}}
	} else if queryBson["decision"] == 1. {
		vote.AgreedUsers = append(vote.AgreedUsers, queryBson["userId"].(string))
		update = bson.M{"$set": bson.M{"agree": vote.Agree + 1, "agreedUsers": vote.AgreedUsers}}
	} else {
		log.Println("Wrong value of decision")
		return nil, nil
	}
	filter := bson.M{"voteId": queryBson["voteId"]}
	res, err := VoteCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Update Vote Error", err)
		return nil, err
	}
	return res, nil
}

func GetVote(db *mongo.Database) ([]*models.Vote, error) {
	VoteCollection := db.Collection("Vote")
	var votes []*models.Vote
	cur, err := VoteCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("GetVote Error", err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Vote Error", err)
			return nil, err
		}
		votes = append(votes, &result)
	}
	return votes, nil
}

// TODO: Check whether the Vote exists
func LaunchVote(db *mongo.Database, query map[string]string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get voteId and update system info
	SystemCollection := db.Collection("System")
	var system models.System
	result := SystemCollection.FindOne(context.Background(), bson.M{})
	sysErr := result.Decode(&system)
	if sysErr != nil {
		log.Println("Get System Error", sysErr)
		return "", sysErr
	}
	voteId := system.TotalVotes
	update := bson.M{"$set": bson.M{"totalVotes": voteId + 1}}
	filter := bson.M{"totalVotes": voteId}
	_, sysErr2 := SystemCollection.UpdateOne(ctx, filter, update)
	if sysErr2 != nil {
		log.Println("Update System Error", sysErr2)
		return "", sysErr2
	}

	// Save Vote into database
	VoteCollection := db.Collection("Vote")
	var vote = models.Vote{
		VoteId:    "vote" + strconv.Itoa(voteId),
		Launcher:  query["user_id"],
		BoardName: query["board_name"],
		Img:       query["img_url"],
		Reason:    query["reason"],
	}
	_, voteErr := VoteCollection.InsertOne(ctx, vote)
	if voteErr != nil {
		log.Println("Insert Vote Error", voteErr)
		return "", voteErr
	}

	// Add voteId to User
	user, userErr := GetUserInfoById(db, models.User{UserId: query["user_id"]})
	if userErr != nil {
		log.Println("Get User Error", userErr)
		return "", userErr
	}
	UserCollection := db.Collection("User")
	filter = bson.M{"$set": bson.M{"userId": query["user_id"]}}
	update = bson.M{"launchNewBoard": append(user.LaunchNewBoard, vote)}
	_, userErr2 := UserCollection.UpdateOne(ctx, filter, update)
	if sysErr2 != nil {
		log.Println("Update User Error", userErr2)
		return "", userErr2
	}

	return vote.VoteId, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
