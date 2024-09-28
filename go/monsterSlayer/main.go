package main

import interaction "aliseylaneh/monsterSlayer/interaction"

var currentRound int = 0

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}
	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)
	interaction.GetPlayerChoice(isSpecialRound)
	return ""
}
func endGame() {

}
