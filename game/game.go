package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}
type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	//use select to process input in channels
	// Print to screen
	// Keep track ofround numbet
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			// changes here to send info back when done
			g.DisplayChan <- ""
		}

	}
}

func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func (g *Game) PrintIntro() {
	g.DisplayChan <- fmt.Sprintf(`
	Rock, Paper & Scissors
	----------------------"
	Game is played for three rounds, and best of three wins the game. Good luck!

	`)
	<-g.DisplayChan
}
func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("----------")

	fmt.Print("Please enter rock, paper, scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.TrimSpace(playerChoice)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
		break
	case PAPER:
		fmt.Println("Computer chose PAPER")
		break
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
		break
	default:
	}
	fmt.Println()

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			// *** decrement i by 1, since we're repeating the round
			g.DisplayChan <- "Invalid choice!"
			return false
		}
	}
	return true

}
func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"

}
func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"

}
func (g *Game) PrintSummary() {
	fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.ComputerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}
}
