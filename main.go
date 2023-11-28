package main

import (
	"myapp/game"
)

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}
	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	// *** added the for loop
	for {
		game.RoundChan <- 1
		<-game.RoundChan
		if game.Round.RoundNumber > 3 {
			break
		}
		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	/* fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
	fmt.Println()
	if playerScore > computerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	} */
}
