package main

import (
	actions "aliseylaneh/monsterSlayer/actions"
	interaction "aliseylaneh/monsterSlayer/interaction"
)

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
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	if userChoice == "ATTACK" {
		actions.AttackMonster(isSpecialRound)
	} else if userChoice == "HEAL" {

	}
	return ""
}
func endGame() {

}
