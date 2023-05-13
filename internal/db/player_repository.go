package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"tv/quick-bat/internal/domain"
)

type PlayerRecord struct {
	Id     int `bson:"_id"`
	Losses int `bson:"losses"`
	Wins   int `bson:"wins"`
}

func UpdatePlayer(player *domain.Player) {
	filter := bson.D{{"_id", player.Id()}}
	update := bson.D{{"$set", bson.D{
		{"wins", player.Wins()},
		{"losses", player.Losses()},
	}}}
	_, err := client.Database("quickbat").Collection("players").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func FetchAllPlayers() []*domain.Player {
	cursor, err := client.Database("quickbat").Collection("players").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	var playerRecords []PlayerRecord
	if err = cursor.All(context.TODO(), &playerRecords); err != nil {
		panic(err)
	}
	var players []*domain.Player
	for _, playerRecord := range playerRecords {
		players = append(players, domain.CreatePlayer(
			domain.PlayerId(playerRecord.Id),
			playerRecord.Wins,
			playerRecord.Losses),
		)
	}

	return players
}
