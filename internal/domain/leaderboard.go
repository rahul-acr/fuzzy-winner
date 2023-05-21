package domain

import "sort"

type LeaderBoard struct {
	players []*Player
}

var MainLeaderBoard *LeaderBoard

func GetLeaderBoard() *LeaderBoard {
	return MainLeaderBoard
}

func (board *LeaderBoard) Len() int {
	return len(board.players)
}

func (board *LeaderBoard) Less(i, j int) bool {
	return board.players[i].wins > board.players[j].wins
}

func (board *LeaderBoard) Swap(i, j int) {
	board.players[i], board.players[j] = board.players[j], board.players[i]
}

func NewLeaderBoard(players []*Player) *LeaderBoard {
	board := LeaderBoard{players: players}
	board.refresh()
	return &board
}

func (board *LeaderBoard) refresh() {
	sort.Sort(board)
}

func (board *LeaderBoard) GetRank(player *Player) int {
	for i := 0; i < board.Len(); i++ {
		if player.id == board.players[i].id {
			return i + 1
		}
	}
	return -1
}

func (board *LeaderBoard) FindPlayer(playerId PlayerId) *Player {
	for i := 0; i < board.Len(); i++ {
		if playerId == board.players[i].id {
			return board.players[i]
		}
	}
	return nil
}
