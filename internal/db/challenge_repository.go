package db

import (
	"context"
	"time"
	"tv/quick-bat/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChallengeRepository struct {
	collection       *mongo.Collection
	playerRepository *PlayerRepository
}

func NewChallengeRepository(
	collection *mongo.Collection,
	playerRepository *PlayerRepository,
) *ChallengeRepository {
	return &ChallengeRepository{collection: collection, playerRepository: playerRepository}
}

func (c *ChallengeRepository) Update(challenge domain.Challenge) {
	update := bson.D{
		{Key: "isAccepted", Value: challenge.IsAccepted()},
		{Key: "time", Value: challenge.Time()},
	}
	winner := challenge.Winner()
	if (winner != domain.Player{}) {
		update = append(update, bson.E{Key: "winner", Value: winner.Id()})
	}
	_, err := c.collection.UpdateByID(
		context.TODO(),
		challenge.GetId(),
		bson.D{{Key: "$set", Value: update}},
	)
	if err != nil {
		panic(err)
	}
}

func (c *ChallengeRepository) Add(challenge domain.Challenge) {
	opponent := challenge.Opponent()
	challenger := challenge.Challenger()

	result, err := c.collection.InsertOne(context.TODO(), ChallengeRecord{
		Id:           primitive.NewObjectID(),
		ChallengerId: int(challenger.Id()),
		OpponentId:   int(opponent.Id()),
	})
	if err != nil {
		panic(err)
	}
	challenge.SetId(result.InsertedID)
}


func (c *ChallengeRepository) FindChallenge(challengeId any) (ChallengeRecord, error) {
	hex := challengeId.(string)
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return ChallengeRecord{}, err
	}
	var record ChallengeRecord
	err = c.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return ChallengeRecord{}, err
	}
	return record, nil
}

type ChallengeRecord struct {
	Id           primitive.ObjectID `bson:"_id"`
	OpponentId   int                `bson:"opponentId"`
	ChallengerId int                `bson:"challengerId"`
	WinnerId     int                `bson:"winnerId"`
	IsAccepted   bool               `bson:"isAccepted"`
	Time         *time.Time         `bson:"time"`
}
