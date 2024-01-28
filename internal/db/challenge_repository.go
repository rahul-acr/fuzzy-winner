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
	return &ChallengeRepository{
		collection:       collection,
		playerRepository: playerRepository,
	}
}

func (c *ChallengeRepository) Update(challenge domain.Challenge) {
	update := bson.D{
		{Key: "isAccepted", Value: challenge.IsAccepted()},
		{Key: "time", Value: challenge.Time()},
	}
	winner := challenge.Winner()
	if (winner != domain.Player{}) {
		update = append(update, bson.E{Key: "winnerId", Value: winner.Id()})
	}
	_, err := c.collection.UpdateByID(
		context.TODO(),
		challenge.Id,
		bson.D{{Key: "$set", Value: update}},
	)
	if err != nil {
		panic(err)
	}
}

func (c *ChallengeRepository) Add(ctx context.Context, challenge domain.Challenge) (domain.Challenge, error) {
	opponent := challenge.Opponent()
	challenger := challenge.Challenger()

	result, err := c.collection.InsertOne(ctx, ChallengeRecord{
		Id:           primitive.NewObjectID(),
		ChallengerId: int(challenger.Id()),
		OpponentId:   int(opponent.Id()),
	})
	if err != nil {
		return domain.Challenge{}, nil
	}
	challenge.Id = result.InsertedID
	return challenge, nil
}

func (c *ChallengeRepository) FindChallenge(ctx context.Context, challengeId any) (domain.Challenge, error) {
	hex := challengeId.(string)
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return domain.Challenge{}, err
	}
	var record ChallengeRecord
	err = c.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return domain.Challenge{}, err
	}
	return c.challengeFromRecord(ctx, record)
}

func (c *ChallengeRepository) FindChallengesForPlayer(ctx context.Context, playerId any) ([]domain.Challenge, error) {
	cursor, err := c.collection.Find(ctx, bson.M{"opponentId": playerId.(int)})
	if err != nil {
		return nil, err
	}

	var records []ChallengeRecord
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	var challenges []domain.Challenge
	for _, r := range records {
		challenge, err := c.challengeFromRecord(ctx, r)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

func (c *ChallengeRepository) FindChallengesByPlayer(ctx context.Context, playerId any) ([]domain.Challenge, error) {
	cursor, err := c.collection.Find(ctx, bson.M{"challengerId": playerId.(int)})
	if err != nil {
		return nil, err
	}

	var records []ChallengeRecord
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	var challenges []domain.Challenge
	for _, r := range records {
		challenge, err := c.challengeFromRecord(ctx, r)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

func (c *ChallengeRepository) challengeFromRecord(ctx context.Context, record ChallengeRecord) (domain.Challenge, error) {
	challenger, err := c.playerRepository.FindPlayer(ctx, record.ChallengerId)
	if err != nil {
		return domain.Challenge{}, err
	}
	opponent, err := c.playerRepository.FindPlayer(ctx, record.OpponentId)
	if err != nil {
		return domain.Challenge{}, err
	}
	var winner domain.Player
	if record.WinnerId != 0 {
		winner, err = c.playerRepository.FindPlayer(ctx, record.WinnerId)
		if err != nil {
			return domain.Challenge{}, err
		}
	}
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
