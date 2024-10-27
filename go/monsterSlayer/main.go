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
	var playerHealValue int
	var monsterAttackDmg int
	var playerAttackDmg int
	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else if userChoice == "SPECIAL_ATTACK" {
		playerAttackDmg = actions.AttackMonster(true)
	} else {
		interaction.InvalidAction()
		return ""
	}
	monsterAttackDmg = actions.AttackPlayer()
	monsterHealth, playerHealth := actions.GetHealthAmount()
	roundData := interaction.NewRoundData(userChoice,
		playerAttackDmg,
		playerHealValue,
		monsterAttackDmg,
		playerHealth,
		monsterHealth)
	interaction.PrintingRoundStatistic(roundData)
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
