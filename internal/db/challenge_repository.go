package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"tv/quick-bat/internal/domain"
)

type ChallengeRepository struct {
	collection *mongo.Collection
}

func NewChallengeRepository(collection *mongo.Collection) *ChallengeRepository {
	return &ChallengeRepository{collection: collection}
}

func (c *ChallengeRepository) Update(challenge *domain.Challenge) {
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

func (c *ChallengeRepository) Add(challenge *domain.Challenge) {
	result, err := c.collection.InsertOne(context.TODO(), struct {
		Challenger int
		Opponent   int
	}{
		Challenger: int(challenge.Challenger().Id()),
		Opponent:   int(challenge.Opponent().Id()),
	})
	if err != nil {
		panic(err)
	}
	challenge.Id = result.InsertedID
}
