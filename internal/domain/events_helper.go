package domain

import "tv/quick-bat/internal/events"

func publishPlayerUpdate(player Player) {
	events.Publish("playerUpdate", player)
}

func publishChallengeCreate(challenge Challenge) {
	events.Publish("challengeCreate", challenge)
}

func publishChallengeUpdate(challenge Challenge) {
	events.Publish("challengeUpdate", challenge)
}

func AddPlayerUpdateListener(listener func(player Player)) {
	events.Listen("playerUpdate", func(event events.Event) {
		listener(event.Payload.(Player))
	})
}

func AddChallengeCreateListener(listener func(challenge Challenge)) {
	events.Listen("challengeCreate", func(event events.Event) {
		listener(event.Payload.(Challenge))
	})
}

func AddChallengeChangeListener(listener func(player Challenge)) {
	events.Listen("challengeUpdate", func(event events.Event) {
		listener(event.Payload.(Challenge))
	})
}
