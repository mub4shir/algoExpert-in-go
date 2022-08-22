package main

import "fmt"

const homeTeamWon = 1

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Println(tournamentWinner(competitions, results))
}

// O(n) time | O(k) space, where n is number of competitions and k is the number of teams competing in these competitions
func tournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}
	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]
		winningTeam := awayTeam
		if result == homeTeamWon {
			winningTeam = homeTeam
		}
		updateScores(winningTeam, 3, scores)
		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}
	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}
