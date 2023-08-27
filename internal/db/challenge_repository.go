package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"tv/quick-bat/internal/domain"
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
		{"isAccepted", challenge.IsAccepted()},
		{"time", challenge.Time()},
	}
	if challenge.Winner() != nil {
		update = append(update, bson.E{Key: "winner", Value: challenge.Winner().Id()})
	}
	_, err := c.collection.UpdateByID(
		context.TODO(),
		challenge.Id,
		bson.D{{"$set", update}},
	)
	if err != nil {
		panic(err)
	}
}

func (c *ChallengeRepository) Add(challenge domain.Challenge) {
	result, err := c.collection.InsertOne(context.TODO(), ChallengeRecord{
		Id:           primitive.NewObjectID(),
		ChallengerId: int(challenge.Challenger().Id()),
		OpponentId:   int(challenge.Opponent().Id()),
	})
	if err != nil {
		panic(err)
	}
	challenge.Id = result.InsertedID
}

func (c *ChallengeRepository) Find(challengeId interface{}) (*domain.Challenge, error) {
	hex := challengeId.(string)
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}
	var record ChallengeRecord
	err = c.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return nil, err
	}
	challenger, err := c.playerRepository.FindById(record.ChallengerId)
	if err != nil {
		return nil, err
	}
	opponent, err := c.playerRepository.FindById(record.OpponentId)
	if err != nil {
		return nil, err
	}
	var winner *domain.Player
	if record.WinnerId != 0 {
		winner, err = c.playerRepository.FindById(record.WinnerId)
		if err != nil {
			return nil, err
		}
	}
	//matchTime, err := time.Parse(time.RFC3339, record.Time)
	//if err != nil {
	//	return nil, err
	//}
	return domain.LoadChallenge(record.Id, challenger, opponent, winner, record.IsAccepted, record.Time), nil
}

type ChallengeRecord struct {
	Id           primitive.ObjectID `bson:"_id"`
	OpponentId   int                `bson:"opponentId"`
	ChallengerId int                `bson:"challengerId"`
	WinnerId     int                `bson:"winnerId"`
	IsAccepted   bool               `bson:"isAccepted"`
	Time         *time.Time         `bson:"time"`
}
