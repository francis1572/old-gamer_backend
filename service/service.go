package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"final_backend/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	UpdateAllVotesStatus(db)

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
	log.Println("task.ToQueryBson()", task.ToQueryBson())
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

		comments, err := GetCommentsByFloor(db, models.Comment{PostId: result.PostId, Floor: result.Floor})
		if err == nil {
			for _, comment := range comments {
				var temp = models.Comment{
					CommentId:  comment.CommentId,
					PostId:     comment.PostId,
					Tag:        comment.Tag,
					Floor:      comment.Floor,
					Content:    comment.Content,
					Author:     comment.Author,
					AuthorName: comment.AuthorName,
					LikeNum:    comment.LikeNum,
					LikedUsers: comment.LikedUsers,
					Time:       comment.Time,
				}
				result.Comments = append(result.Comments, temp)
			}
		}
		log.Println("有通過第六階段", result.Comments)

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

func GetCommentsByFloor(db *mongo.Database, task models.Comment) ([]*models.Comment, error) {
	CommentCollection := db.Collection("Comment")
	var comments []*models.Comment
	cur, err := CommentCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Comment{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		comments = append(comments, &result)
	}
	return comments, nil
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

func InsertPost(db *mongo.Database, task models.Post) (*mongo.InsertManyResult, error) {
	PostCollection := db.Collection("Post")
	var post = models.PostDB{
		PostId:       task.PostId,
		BoardId:      task.BoardId,
		ChildBoardId: task.ChildBoardId,
		PostTag:      task.PostTag,
		PostTitle:    task.PostTitle,
		Author:       task.Author,
		AuthorName:   task.AuthorName,
		Floor:        task.Floor,
		CommentNum:   0,
		LikeNum:      0,
		Time:         time.Now(),
		LikedUsers:   make([]string, 0),
	}
	postId := uuid.New().String()
	if post.PostId == "" {
		post.PostId = postId
		for i := range task.Content {
			task.Content[i].PostId = postId
		}
	}

	BlockCollection := db.Collection("Block")
	BlockList := make([]interface{}, len(task.Content))
	for i := range task.Content {
		BlockList[i] = task.Content[i]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := PostCollection.InsertOne(ctx, post)
	if err != nil {
		log.Println("Insert post Error", err)
		return nil, err
	}
	res, err := BlockCollection.InsertMany(ctx, BlockList)
	if err != nil {
		log.Println("Insert blocks Error", err)
		return nil, err
	}

	return res, nil
}

func UpdatePost(db *mongo.Database, task models.Post) error {
	PostCollection := db.Collection("Post")
	var post = models.PostDB{
		PostId:       task.PostId,
		BoardId:      task.BoardId,
		ChildBoardId: task.ChildBoardId,
		PostTag:      task.PostTag,
		PostTitle:    task.PostTitle,
		Author:       task.Author,
		AuthorName:   task.AuthorName,
		Floor:        task.Floor,
		CommentNum:   task.CommentNum,
		LikeNum:      task.LikeNum,
		Time:         task.Time,
		LikedUsers:   task.LikedUsers,
	}

	BlockCollection := db.Collection("Block")
	BlockList := make([]models.Block, len(task.Content))
	for i := range task.Content {
		BlockList[i] = task.Content[i]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"postId": post.PostId, "floor": post.Floor}
	update := bson.M{"$set": bson.M{"postTag": task.PostTag, "postTitle": task.PostTitle}}
	_, err := PostCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Insert post Error", err)
		return err
	}

	for _, block := range BlockList {
		opts := options.Update().SetUpsert(true)
		filter = bson.M{"postId": block.PostId, "floor": block.Floor, "blockId": block.BlockId}
		update = bson.M{"$set": bson.M{"subtitle": block.Subtitle, "content": block.Content}}
		_, err = BlockCollection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Println("Insert post Error", err)
			return err
		}
	}

	return nil
}

func DeletePost(db *mongo.Database, task models.Post) error {
	// PostCollection := db.Collection("Post")
	BlockCollection := db.Collection("Block")
	var emptyBlock = models.Block{
		PostId:   task.PostId,
		Floor:    task.Floor,
		BlockId:  0,
		Subtitle: "此篇文章已被刪除。",
		Content:  "此篇文章已被刪除。",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// _, err := PostCollection.DeleteOne(ctx, bson.M{"postId": task.PostId, "floor": task.Floor})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	_, err := BlockCollection.DeleteMany(ctx, bson.M{"postId": task.PostId, "floor": task.Floor})
	if err != nil {
		log.Fatal(err)
	}
	_, err = BlockCollection.InsertOne(ctx, emptyBlock)
	if err != nil {
		log.Println("Insert post Error", err)
	}

	return nil
}

func DeleteComment(db *mongo.Database, task models.Comment) error {
	CommentCollection := db.Collection("Comment")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := CommentCollection.DeleteOne(ctx, bson.M{"commentId": task.CommentId})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteCitation(db *mongo.Database, task models.Citation) error {
	CitationCollection := db.Collection("Citation")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := CitationCollection.DeleteOne(ctx, bson.M{"citationId": task.CitationId})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func InsertComment(db *mongo.Database, task models.Comment) (*mongo.InsertOneResult, error) {
	CommentCollection := db.Collection("Comment")
	var comment = models.Comment{
		CommentId:  uuid.New().String(),
		PostId:     task.PostId,
		Tag:        task.Tag,
		Floor:      task.Floor,
		Content:    task.Content,
		Author:     task.Author,
		AuthorName: task.AuthorName,
		LikeNum:    0,
		LikedUsers: make([]string, 0),
		Time:       time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := CommentCollection.InsertOne(ctx, comment)
	if err != nil {
		log.Println("Insert post Error", err)
		return nil, err
	}

	return res, nil
}

func UpdateComment(db *mongo.Database, task models.Comment) error {
	CommentCollection := db.Collection("Comment")
	var comment = models.Comment{
		CommentId:  task.CommentId,
		PostId:     task.PostId,
		Tag:        task.Tag,
		Floor:      task.Floor,
		Content:    task.Content,
		Author:     task.Author,
		AuthorName: task.AuthorName,
		LikeNum:    task.LikeNum,
		LikedUsers: task.LikedUsers,
		Time:       task.Time,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"commentId": comment.CommentId}
	update := bson.M{"$set": bson.M{"tag": comment.Tag, "content": comment.Content}}
	_, err := CommentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Insert comment Error", err)
		return err
	}

	return nil
}

func UpdateUser(db *mongo.Database, task models.User) error {
	UserCollection := db.Collection("GUser")
	var user = models.User{
		UserId:           task.UserId,
		SelfIntroduction: task.SelfIntroduction,
		InterestGames:    task.InterestGames,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": user.UserId}
	update := bson.M{"$set": bson.M{"selfIntroduction": user.SelfIntroduction, "interestGames": user.InterestGames}}
	_, err := UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Insert comment Error", err)
		return err
	}

	return nil
}

func LikePost(db *mongo.Database, task models.Post) error {
	// newID := uuid.New().String()
	// log.Println(newID)
	PostCollection := db.Collection("Post")
	var post models.Post
	result := PostCollection.FindOne(context.Background(), task.ToQueryBson())
	err := result.Decode(&post)
	post.LikeNum = post.LikeNum + 1
	post.LikedUsers = append(post.LikedUsers, task.Author)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"postId": post.PostId, "floor": task.Floor}
	update := bson.M{"$set": bson.M{"likeNum": post.LikeNum, "likedUsers": post.LikedUsers}}
	_, err = PostCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println("Update Comment Error", err)
		return err
	}
	return err
}

func LikeComment(db *mongo.Database, task models.Comment) error {
	CommentCollection := db.Collection("Comment")
	var comment models.Comment
	result := CommentCollection.FindOne(context.Background(), task.ToQueryBson())
	err := result.Decode(&comment)
	comment.LikeNum = comment.LikeNum + 1
	comment.LikedUsers = append(comment.LikedUsers, task.Author)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"commentId": task.CommentId}
	update := bson.M{"$set": bson.M{"likeNum": comment.LikeNum, "likedUsers": comment.LikedUsers}}
	_, err = CommentCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println("Update Comment Error", err)
		return err
	}
	return err
}

func InsertCitation(db *mongo.Database, task models.Citation) (*mongo.InsertOneResult, error) {
	CitationCollection := db.Collection("Citation")
	var citation = models.Citation{
		CitationId: uuid.New().String(),
		PostId:     task.PostId,
		Floor:      task.Floor,
		CitedFloor: task.CitedFloor,
		BlockId:    task.BlockId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := CitationCollection.InsertOne(ctx, citation)
	if err != nil {
		log.Println("Insert Citation Error", err)
		return nil, err
	}

	return res, nil
}

func GetVoteById(db *mongo.Database, query models.Vote) (*models.Vote, error) {
	UpdateAllVotesStatus(db)

	VoteCollection := db.Collection("Vote")
	var vote models.Vote
	result := VoteCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&vote)
	if err != nil {
		log.Println("Decode vote Error", err)
		return nil, err
	}

	user, _ := GetUserInfoById(db, models.User{UserId: vote.Launcher})
	vote.LauncherInfo = *user

	return &vote, nil
}

func Vote(db *mongo.Database, queryBson bson.M) (*mongo.UpdateResult, error) {
	UpdateAllVotesStatus(db)

	VoteCollection := db.Collection("Vote")
	var vote models.Vote
	result := VoteCollection.FindOne(context.Background(), bson.M{"voteId": queryBson["voteId"]})
	err := result.Decode(&vote)
	if err != nil {
		log.Println("Decode vote Error", err)
		return nil, err
	}

	SystemCollection := db.Collection("System")
	var system models.System
	result = SystemCollection.FindOne(context.Background(), bson.M{})
	err = result.Decode(&system)
	if err != nil {
		log.Println("Get System Error", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the user voted
	if contains(vote.DisagreedUsers, queryBson["userId"].(string)) || contains(vote.AgreedUsers, queryBson["userId"].(string)) {
		log.Println("User voted")
		return nil, nil
	}

	// Check if the vote ended
	if vote.Status != "active" {
		log.Println("Vote ended")
		return nil, nil
	}

	var update bson.M
	status := vote.Status
	// type of queryBson["decision"] is float
	if queryBson["decision"] == 0. {
		update = bson.M{"$set": bson.M{
			"disagree":       vote.Disagree + 1,
			"disagreedUsers": append(vote.DisagreedUsers, queryBson["userId"].(string)),
		}}
	} else if queryBson["decision"] == 1. {
		if vote.Agree+1 == system.VoteThreshold {
			status = "success"
			LaunchBoard(db, vote)
		}
		update = bson.M{"$set": bson.M{
			"agree":       vote.Agree + 1,
			"agreedUsers": append(vote.AgreedUsers, queryBson["userId"].(string)),
			"status":      status,
		}}
	} else {
		log.Println("Wrong value of decision")
		return nil, nil
	}
	filter := bson.M{"voteId": queryBson["voteId"]}
	res, err := VoteCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Vote Error", err)
		return nil, err
	}
	return res, nil
}

func GetVote(db *mongo.Database) ([]models.Vote, error) {
	UpdateAllVotesStatus(db)

	VoteCollection := db.Collection("Vote")
	var votes []models.Vote
	cur, err := VoteCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("GetVote Error", err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		// log.Println("!!!", result.VoteId, result.Status)
		if err != nil {
			log.Println("Decode Vote Error", err)
			return nil, err
		}
		votes = append(votes, result)
	}
	return votes, nil
}

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
		VoteId:         "vote" + strconv.Itoa(voteId),
		Launcher:       query["userId"],
		BoardName:      query["boardName"],
		DomainName:     query["domainName"],
		Img:            query["imgUrl"],
		Reason:         query["reason"],
		Agree:          1,
		Disagree:       0,
		AgreedUsers:    []string{query["userId"]},
		DisagreedUsers: make([]string, 0),
		Deadline:       time.Now().AddDate(0, 1, 0),
		Status:         "active",
	}
	_, voteErr := VoteCollection.InsertOne(ctx, vote)
	if voteErr != nil {
		log.Println("Insert Vote Error", voteErr)
		return "", voteErr
	}

	// Add voteId to User
	user, userErr := GetUserInfoById(db, models.User{UserId: query["userId"]})
	if userErr != nil {
		log.Println("Get User Error", userErr)
		return "", userErr
	}
	UserCollection := db.Collection("User")
	filter = bson.M{"$set": bson.M{"userId": query["userId"]}}
	update = bson.M{"launchNewBoard": append(user.LaunchNewBoard, vote)}
	_, userErr2 := UserCollection.UpdateOne(ctx, filter, update)
	if sysErr2 != nil {
		log.Println("Update User Error", userErr2)
		return "", userErr2
	}

	return vote.VoteId, nil
}

func UpdateAllVotesStatus(db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	VoteCollection := db.Collection("Vote")
	var votes []*models.Vote
	cur, err := VoteCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("GetVote Error", err)
	}

	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Vote Error", err)
		}
		votes = append(votes, &result)
	}

	for _, vote := range votes {
		if vote.Status == "active" {
			if time.Now().After(vote.Deadline) {
				filter := bson.M{"voteId": vote.VoteId}
				update := bson.M{"$set": bson.M{"status": "fail"}}
				_, err := VoteCollection.UpdateOne(ctx, filter, update)
				if err != nil {
					log.Println("Update Vote Status Error", err)
				}
			}
		}
	}
}

func LaunchBoard(db *mongo.Database, vote models.Vote) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	SystemCollection := db.Collection("System")
	var system models.System
	result := SystemCollection.FindOne(context.Background(), bson.M{})
	err := result.Decode(&system)
	if err != nil {
		log.Println("Get System Error", err)
	}
	boardId := system.TotalBoards
	update := bson.M{"$set": bson.M{"totalBoards": system.TotalBoards + 1}}
	filter := bson.M{"totalBoards": system.TotalBoards}
	_, err = SystemCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Update System Error", err)
	}

	board := models.Board{
		BoardId:     "bd" + fmt.Sprint(boardId),
		BoardName:   vote.BoardName,
		DomainName:  vote.DomainName,
		PostNum:     0,
		Img:         vote.Img,
		ChildBoards: make([]models.ChildBoard, 0),
	}
	BoardCollection := db.Collection("Board")
	_, err = BoardCollection.InsertOne(ctx, board)
	if err != nil {
		log.Println("Insert board Error", err)
	}
}

func GetNotification(db *mongo.Database, query models.User) (*models.Notification, error) {
	NotificationCollection := db.Collection("Notification")
	var notification *models.Notification
	result := NotificationCollection.FindOneAndDelete(context.Background(), query.ToQueryBson())
	err := result.Decode(&notification)
	if err != nil {
		log.Println("Decode notification Error", err)
		return nil, err
	}

	return notification, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
