package domain

type PlayerId int

type Match struct {
	firstPlayerId  PlayerId
	secondPlayerId PlayerId
	winnerPlayerId PlayerId
}
