package db

import (
	"context"
	"tv/quick-bat/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchRepository struct {
	collection       *mongo.Collection
	playerRepository *PlayerRepository
}

type MatchRecord struct {
	Id       primitive.ObjectID `bson:"_id"`
	WinnerId int                `bson:"winnerId"`
	LoserId  int                `bson:"loserId"`
}

func NewMatchRepository(collection *mongo.Collection, playerRepository *PlayerRepository) MatchRepository {
	return MatchRepository{collection: collection, playerRepository: playerRepository}
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

func (c MatchRepository) FindMatchesOf(ctx context.Context, playerId domain.PlayerId) ([]domain.Match, error) {
	cursor, err := c.collection.Find(ctx, bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "winnerId", Value: int(playerId)}},
			bson.D{{Key: "loserId", Value: int(playerId)}},
		}},
	})
	var records []MatchRecord
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	matches := make([]domain.Match, len(records))
	for i, r := range records {
		match, err := c.recordToMatch(ctx, r)
		if err != nil {
			return nil, err
		}
		matches[i] = match
	}
	return matches, nil
}

func (c MatchRepository) recordToMatch(ctx context.Context, record MatchRecord) (domain.Match, error) {
	winner, err := c.playerRepository.FindPlayer(ctx, record.WinnerId)
	if err != nil {
		return domain.Match{}, nil
	}
	loser, err := c.playerRepository.FindPlayer(ctx, record.LoserId)
	if err != nil {
		return domain.Match{}, nil
	}
	return domain.Match{
		Id:     record.Id,
		Winner: winner,
		Loser:  loser,
	}, nil
}
