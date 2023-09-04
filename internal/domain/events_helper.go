package domain

import "tv/quick-bat/internal/events"

func publishPlayerUpdate(player Player) {
	events.Publish("playerUpdate", player)
}

func publishChallengeUpdate(challenge Challenge) {
	events.Publish("challengeUpdate", challenge)
}

func AddPlayerUpdateListener(listener func(player Player)) {
	events.Listen("playerUpdate", func(event events.Event) {
		listener(event.Payload.(Player))
	})
}

func AddChallengeChangeListener(listener func(player Challenge)) {
	events.Listen("challengeUpdate", func(event events.Event) {
		listener(event.Payload.(Challenge))
	})
}
