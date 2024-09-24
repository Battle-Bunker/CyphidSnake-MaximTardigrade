package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
	"math"
)

// HeuristicFloodFill calculates a heuristic score based on the available space around the snake
func HeuristicFloodFill(snapshot agent.GameSnapshot) float64 {
	const (
		minSpaceMultiplier = 2 // Minimum desired space multiplier relative to snake length
		maxScore           = 100.0
	)

	totalScore := 0.0

	for _, allySnake := range snapshot.YourTeam() {
		head := allySnake.Head()
		snakeLength := allySnake.Length()
		minDesiredSpace := snakeLength * minSpaceMultiplier

		availableSpace := floodFill(snapshot, head)

		// Calculate score based on available space
		if availableSpace >= minDesiredSpace {
			totalScore += maxScore
		} else {
			// Partial score if the space is less than desired
			totalScore += (float64(availableSpace) / float64(minDesiredSpace)) * maxScore
		}
	}

	output := totalScore / float64(len(snapshot.YourTeam()))

	if math.IsNaN(output) {
		return 0
	} else {
		return output
	}
	// Average the score across all team snakes

}

// floodFill performs a flood fill from the given point and returns the number of accessible cells
func floodFill(snapshot agent.GameSnapshot, start rules.Point) int {
	visited := make(map[rules.Point]bool)
	queue := []rules.Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		// Check adjacent cells
		for _, move := range []string{"up", "down", "left", "right"} {
			next := getAdjacentPoint(current, move)
			if isValidMove(snapshot, next) && !visited[next] {
				queue = append(queue, next)
			}
		}
	}

	return len(visited)
}

// getAdjacentPoint returns the adjacent point based on the given move
func getAdjacentPoint(p rules.Point, move string) rules.Point {
	switch move {
	case "up":
		return rules.Point{X: p.X, Y: p.Y + 1}
	case "down":
		return rules.Point{X: p.X, Y: p.Y - 1}
	case "left":
		return rules.Point{X: p.X - 1, Y: p.Y}
	case "right":
		return rules.Point{X: p.X + 1, Y: p.Y}
	default:
		return p
	}
}

// isValidMove checks if the given point is a valid move on the board
func isValidMove(snapshot agent.GameSnapshot, p rules.Point) bool {
	// Check if the point is within the board boundaries
	if p.X < 0 || p.X >= snapshot.Width() || p.Y < 0 || p.Y >= snapshot.Height() {
		return false
	}

	// Check if the point collides with any snake body
	for _, snake := range snapshot.Snakes() {
		for _, bodyPart := range snake.Body() {
			if p == bodyPart {
				return false
			}
		}
	}

	// Check if the point is in a hazard (you may want to allow hazards but with a penalty)
	for _, hazard := range snapshot.Hazards() {
		if p == hazard {
			return false
		}
	}

	return true
}
