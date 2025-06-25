// "Футбольный анализ"

package main

import (
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
		Rating:  calculateRating(goals, assists, misses),
	}
}

func calculateRating(goals, assists, misses int) float64 {
	if misses == 0 {
		return (float64(goals) + float64(assists)/2)
	}
	return (float64(goals) + float64(assists)/2) / float64(misses)
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals > players[j].Goals
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		var ratioA, ratioB float64
		if players[i].Misses == 0 {
			ratioA = float64(players[i].Goals + players[i].Assists/2)
		} else {
			ratioA = float64(players[i].Goals+players[i].Assists/2) / float64(players[i].Misses)
		}

		if players[j].Misses == 0 {
			ratioB = float64((players[j].Goals + players[j].Assists) / 2)
		} else {
			ratioB = float64(players[j].Goals+players[j].Assists/2) / float64(players[j].Misses)
		}
		if ratioA == ratioB {
			return players[i].Name < players[j].Name
		}
		return ratioA > ratioB
	})
	return players
}
