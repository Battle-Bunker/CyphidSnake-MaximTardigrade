package main

import (
    "github.com/BattlesnakeOfficial/rules"
    "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicSpaceControl calculates a score based on the amount of space
// controlled by your team's snakes compared to opponents' snakes
func HeuristicSpaceControl(snapshot agent.GameSnapshot) float64 {
    var score float64

    // Calculate controlled space for your team
    for _, allySnake := range snapshot.YourTeam() {
        if allySnake.Alive() {
            score += float64(floodFill(snapshot, allySnake.Head()))
        }
    }

    // Subtract controlled space for opponents
    for _, opponentSnake := range snapshot.Opponents() {
        if opponentSnake.Alive() {
            score -= float64(floodFill(snapshot, opponentSnake.Head()))
        }
    }

    return score
}

// floodFill performs a flood fill from the given point and returns the number of accessible cells
func floodFill(snapshot agent.GameSnapshot, start rules.Point) int {
    visited := make(map[rules.Point]bool)
    queue := []rules.Point{start}
    count := 0

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if visited[current] {
            continue
        }

        visited[current] = true
        count++

        // Check adjacent cells
        directions := []rules.Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
        for _, dir := range directions {
            next := rules.Point{X: current.X + dir.X, Y: current.Y + dir.Y}
            if isValidMove(snapshot, next) && !visited[next] {
                queue = append(queue, next)
            }
        }
    }

    return count
}

// isValidMove checks if a move to the given point is valid
func isValidMove(snapshot agent.GameSnapshot, point rules.Point) bool {
    // Check if the point is within the board
    if point.X < 0 || point.X >= snapshot.Width() || point.Y < 0 || point.Y >= snapshot.Height() {
        return false
    }

    // Check if the point collides with any snake
    for _, snake := range snapshot.AllSnakes() {
        for _, bodyPart := range snake.Body() {
            if bodyPart == point {
                return false
            }
        }
    }

    // Check if the point is a hazard
    for _, hazard := range snapshot.Hazards() {
        if hazard == point {
            return false
        }
    }

    return true
}