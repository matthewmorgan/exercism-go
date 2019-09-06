package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Stat struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

var OUTCOMES = []string{"win", "loss", "draw"}
var stats map[string]*Stat

func Tally(reader io.Reader, writer io.Writer) error {
	stats = make(map[string]*Stat)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if shouldIgnoreLine(text) {
			continue
		}
		parts := strings.Split(text, ";")
		if len(parts) < 3 {
			return errors.New("too few arguments, expected 3")
		}
		team1 := parts[0]
		team2 := parts[1]
		outcome := parts[2]
		if !inArray(OUTCOMES, outcome) {
			return errors.New(fmt.Sprintf("illegal argument for outcome: %s", outcome))
		}
		maybeInitializeStat(team1)
		maybeInitializeStat(team2)
		recordScore(team1, team2, outcome)
	}
	return writeTable(writer)
}

func shouldIgnoreLine(text string) bool {
	return len(text) == 0 || text[0] == '#' || text[0] == '\n'
}

func inArray(array []string, target string) bool {
	for _, el := range array {
		if el == target {
			return true
		}
	}
	return false
}

func maybeInitializeStat(team string) {
	if _, exists := stats[team]; !exists {
		stats[team] = new(Stat)
		stats[team].name = team
	}
}

func recordScore(team1 string, team2 string, outcome string) {
	stats[team1].matchesPlayed += 1
	stats[team2].matchesPlayed += 1
	switch outcome {
	case "win":
		{
			stats[team1].wins += 1
			stats[team2].losses += 1
			stats[team1].points += 3
			break
		}
	case "loss":
		{
			stats[team2].wins += 1
			stats[team1].losses += 1
			stats[team2].points += 3
			break
		}
	case "draw":
		{
			stats[team1].points += 1
			stats[team1].draws += 1
			stats[team2].points += 1
			stats[team2].draws += 1
			break
		}
	}
}

func writeTable(writer io.Writer) error {
	sortedStats := []*Stat{}
	for _, stat := range stats {
		sortedStats = append(sortedStats, stat)
	}
	// Sort by points desc, then by name asc.
	sort.Slice(sortedStats, func(i, j int) bool {
		if (sortedStats[i].points == sortedStats[j].points) {
			return sortedStats[i].name < sortedStats[j].name
		}
		return sortedStats[i].points > sortedStats[j].points
	})

	lines := []string{}
	lines = append(lines, fmt.Sprintf("%-31v| %2v | %2v | %2v | %2v | %2v", "Team", "MP", "W", "D", "L", "P"))
	for _, team := range sortedStats {
		line := fmt.Sprintf("%-31v| %2v | %2v | %2v | %2v | %2v", team.name, team.matchesPlayed, team.wins, team.draws, team.losses, team.points)
		lines = append(lines, line)
	}
	writer.Write([]byte(strings.Join(lines[:], "\n")))
	writer.Write([]byte("\n"))
	return nil
}
