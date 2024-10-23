package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var monsterHealth = MONSTER_HEALTH
var playerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_DAMAGE
	maxAttackValue := PLAYER_ATTACK_MAX_DAMAGE
	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DAMAGE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DAMAGE
	}
	damageValue := generateRandomNumBetween(minAttackValue, maxAttackValue)
	monsterHealth -= damageValue

}

func Heal() {
	minHealValue := PLAYER_HEAL_MIN_VALUE
	maxHealValue := PLAYER_HEAL_MAX_VALUE
	healValue := generateRandomNumBetween(minHealValue, maxHealValue)
	healthDiff := PLAYER_HEALTH - playerHealth
	if healthDiff >= healValue {
		playerHealth += healValue
	} else {
		playerHealth = 100
	}
	playerHealth += PLAYER_HEALTH
}

func AttackPlayer() {
	minAttackValue := MONSTER_ATTACK_MIN_DAMAGE
	maxAttackValue := MONSTER_ATTACK_MAX_DAMAGE
	damageValue := generateRandomNumBetween(minAttackValue, maxAttackValue)
	playerHealth -= damageValue
}
func generateRandomNumBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
