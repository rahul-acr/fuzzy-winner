package domain

import "tv/quick-bat/internal/events"

func publishPlayerChange(player *Player) {
	events.Publish("playerUpdate", player)
}

func publishChallengeCreate(challenge *Challenge) {
	events.Publish("challengeCreate", challenge)
}

func publishChallengeUpdate(challenge *Challenge) {
	events.Publish("challengeUpdate", challenge)
}
