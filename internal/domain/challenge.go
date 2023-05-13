package domain

import "time"

type Challenge struct {
	challenger *Player
	opponent   *Player
	isAccepted bool
	time       time.Time
}
