package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(fName string, fScore int, sName string, sScore int) {
	if fScore > sScore {
		l.Wins[fName]++
	} else {
		l.Wins[sName]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for _, v := range l.Teams {
		names = append(names, v.Name)
	}
	slices.SortFunc(names, func(a, b string) int {
		return l.Wins[b] - l.Wins[a]
	})
	return names
}

func main() {
	t1 := Team{
		Name:        "Winners",
		PlayerNames: []string{"Bob", "Billy", "Bruiser"},
	}

	t2 := Team{
		Name:        "Losers",
		PlayerNames: []string{"Cam", "Cameron", "Carrot"},
	}
	t3 := Team{
		Name:        "Other Guys",
		PlayerNames: []string{"Dan", "Debbie", "Dom"},
	}
	league := League{
		Teams: []Team{t1, t2, t3},
		Wins:  map[string]int{},
	}
	league.MatchResult("Winners", 100, "Losers", 10)
	league.MatchResult("Winners", 200, "Losers", 20)
	league.MatchResult("Other Guys", 100, "Winners", 10)
	league.MatchResult("Other Guys", 100, "Losers", 10)
	RankPrinter(league, os.Stdout)
}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	for i, v := range ranking {
		io.WriteString(w, fmt.Sprintf("%d - %s\n", i+1, v))
	}
}
