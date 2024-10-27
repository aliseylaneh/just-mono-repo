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
	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	var monsterHealth int
	var playerHealth int
	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}
	actions.AttackPlayer()
	monsterHealth, playerHealth = actions.GetHealthAmount()
	if playerHealth <= 0 {
		return "MONSTER"
	} else if monsterHealth <= 0 {
		return "PLAYER"
	}
	return ""
}
func endGame(winner string) {
	interaction.DeclareWinner(winner)

}
