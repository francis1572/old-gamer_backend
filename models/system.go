package models

type System struct {
	TotalVotes    int   `bson:"totalVotes" json:"totalVotes"`
	VoteThreshold int64 `bson:"voteThreshold" json:"voteThreshold"`
	TotalBoards   int   `bson:"totalBoards" json:"totalBoards"`
}
