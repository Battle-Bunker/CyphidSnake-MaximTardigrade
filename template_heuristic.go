package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)


// TODO implement a heuristic
func HeuristicWhateverYouWantItToDo(snapshot agent.GameSnapshot) float64 {
	// insert code logic here
	
	return 0 // return the score of the position
}

// once done, add to main.go
/*
	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(1.0, "team-health", HeuristicHealth),
		...
*Add following line*
agent.NewHeuristic(1.0, "whatever-you-want-it-to-do", HeuristicWhateverYouWantItToDo),
	)

*/