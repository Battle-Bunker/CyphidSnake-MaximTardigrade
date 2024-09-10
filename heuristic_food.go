package main

import (
	"github.com/BattlesnakeOfficial/rules"
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"math"
)

// HeuristicFoodPriority calculates a heuristic score based on the team's health and proximity to food
func HeuristicFoodPriority(snapshot agent.GameSnapshot) float64 {
	var totalScore float64 = 0

	for _, allySnake := range snapshot.YourTeam() {
		snakeScore := float64(allySnake.Health())

		// Increase priority for food as health decreases
		foodPriority := math.Max(0, 100-float64(allySnake.Health())) / 100

		// Find distance to nearest food
		minFoodDistance := math.Inf(1)
		for _, food := range snapshot.Food() {
			distance := manhattanDistance(allySnake.Head(), food)
			if float64(distance) < minFoodDistance {
				minFoodDistance = float64(distance)
			}
		}

		// Adjust score based on food proximity and priority
		if minFoodDistance != math.Inf(1) {
			foodScore := (100 - minFoodDistance) * foodPriority
			snakeScore += foodScore
		}

		totalScore += snakeScore
	}

	return totalScore
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}