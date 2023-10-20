package db

import (
	"context"
	"tv/quick-bat/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRecord struct {
	Id     int    `bson:"_id"`
	Name   string `bson:"name"`
	Losses int    `bson:"losses"`
	Wins   int    `bson:"wins"`
}

type PlayerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(collection *mongo.Collection) *PlayerRepository {
	return &PlayerRepository{collection: collection}
}

func (r PlayerRepository) Update(player domain.Player) {
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "wins", Value: player.Wins()},
		{Key: "losses", Value: player.Losses()},
	}}}
	_, err := r.collection.UpdateByID(context.TODO(), player.Id(), update)
	if err != nil {
		panic(err)
	}
}

func (r PlayerRepository) FetchAll(ctx context.Context) []domain.Player {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil
	}

	var playerRecords []PlayerRecord
	if err = cursor.All(ctx, &playerRecords); err != nil {
		panic(err)
	}
	var players []domain.Player
	for _, playerRecord := range playerRecords {
		players = append(players, domain.NewPlayer(
			domain.PlayerId(playerRecord.Id),
			playerRecord.Name,
			playerRecord.Wins,
			playerRecord.Losses),
		)
	}

	return players
}

func (r PlayerRepository) FindPlayer(ctx context.Context, id int) (domain.Player, error) {
	var record PlayerRecord
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return domain.Player{}, err
	}
	return domain.NewPlayer(domain.PlayerId(record.Id), record.Name, record.Wins, record.Losses), nil
}
