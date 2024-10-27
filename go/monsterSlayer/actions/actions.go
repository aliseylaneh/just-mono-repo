package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var monsterHealth = MONSTER_HEALTH
var playerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DAMAGE
	maxAttackValue := PLAYER_ATTACK_MAX_DAMAGE
	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DAMAGE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DAMAGE
	}
	damageValue := generateRandomNumBetween(minAttackValue, maxAttackValue)
	monsterHealth -= damageValue
	return damageValue
}

func HealPlayer() int {
	minHealValue := PLAYER_HEAL_MIN_VALUE
	maxHealValue := PLAYER_HEAL_MAX_VALUE
	healValue := generateRandomNumBetween(minHealValue, maxHealValue)
	healthDiff := PLAYER_HEALTH - playerHealth
	if healthDiff >= healValue {
		playerHealth += healValue
		return healValue
	} else {
		playerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DAMAGE
	maxAttackValue := MONSTER_ATTACK_MAX_DAMAGE
	damageValue := generateRandomNumBetween(minAttackValue, maxAttackValue)
	playerHealth -= damageValue
	return damageValue
}

func GetHealthAmount() (int, int) {
	return monsterHealth, playerHealth
}
func generateRandomNumBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
