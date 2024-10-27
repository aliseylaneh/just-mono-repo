package interaction

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func NewRoundData(action string,
	playerAttackDmg int,
	playerHealValue int,
	monsterAttackDmg int,
	playerHealth int,
	monsterHealth int) *RoundData {
	roundData := RoundData{
		Action:           action,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}
	return &roundData
}

func PrintGreeting() {
	fmt.Println("Monster Slayer")
	fmt.Println("Starting new game...")
	fmt.Println("Good luck.")
}

func ShowAvailableActions(specialAttackAvailable bool) {
	fmt.Print("Please choose your action:\n" +
		"1) Attack Monster\n" +
		"2) Heal\n")
	if specialAttackAvailable {
		fmt.Print("3) Special Attack\n:")
	} else {
		fmt.Print(":")
	}
}
func DeclareWinner(winner string) {
	fmt.Println("------------------------------")
	fmt.Println("----------Game Over-----------")
	fmt.Println("------------------------------")
	fmt.Printf("%v won :)\n", winner)
}
func PrintingRoundStatistic(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)

	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player perfromed a strong attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "HEAL" {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v\n", roundData.MonsterHealth)

}
func InvalidAction() {
	fmt.Println("Invalid action.")
}
