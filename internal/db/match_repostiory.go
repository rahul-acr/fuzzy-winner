package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tv/quick-bat/internal/domain"
)

type MatchRepository struct {
	collection *mongo.Collection
}

type MatchRecord struct {
	Id       primitive.ObjectID `bson:"_id"`
	WinnerId int                `bson:"winnerId"`
	LoserId  int                `bson:"loserId"`
}

func NewMatchRepository(collection *mongo.Collection) MatchRepository {
	return MatchRepository{collection: collection}
}

func (c MatchRepository) Add(ctx context.Context, match domain.Match) (domain.Match, error) {
	result, err := c.collection.InsertOne(ctx, MatchRecord{
		Id:       primitive.NewObjectID(),
		WinnerId: int(match.Winner.Id()),
		LoserId:  int(match.Loser.Id()),
	})
	if err != nil {
		return domain.Match{}, nil
	}
	match.Id = result.InsertedID
	return match, nil
}
