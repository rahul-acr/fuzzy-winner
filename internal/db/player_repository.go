package db

import (
	"context"
	"tv/quick-bat/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRecord struct {
	Id     int `bson:"_id"`
	Losses int `bson:"losses"`
	Wins   int `bson:"wins"`
}

type PlayerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(collection *mongo.Collection) *PlayerRepository {
	return &PlayerRepository{collection: collection}
}

func (r *PlayerRepository) Update(player domain.Player) {
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "wins", Value: player.Wins()},
		{Key: "losses", Value: player.Losses()},
	}}}
	_, err := r.collection.UpdateByID(context.TODO(), player.Id(), update)
	if err != nil {
		panic(err)
	}
}

func (r *PlayerRepository) FetchAll() []domain.Player {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	var playerRecords []PlayerRecord
	if err = cursor.All(context.TODO(), &playerRecords); err != nil {
		panic(err)
	}
	var players []domain.Player
	for _, playerRecord := range playerRecords {
		players = append(players, domain.NewPlayer2(
			domain.PlayerId(playerRecord.Id),
			playerRecord.Wins,
			playerRecord.Losses),
		)
	}

	return players
}

func (r *PlayerRepository) FindById(id int) (*domain.Player, error) {
	//objectId, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	//if err != nil {
	//	return nil, err
	//}
	var playerRecord PlayerRecord
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&playerRecord)
	if err != nil {
		return nil, err
	}
	return domain.NewPlayer(
		domain.PlayerId(playerRecord.Id),
		playerRecord.Wins,
		playerRecord.Losses,
	), nil
}

func (r *PlayerRepository) FindPlayer(id int) (PlayerRecord, error) {
	var playerRecord PlayerRecord
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&playerRecord)
	if err != nil {
		return PlayerRecord{}, err
	}
	return playerRecord, nil
}