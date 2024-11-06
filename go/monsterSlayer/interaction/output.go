package interaction

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"path/filepath"
)

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
	myAscii := figure.NewFigure("Monster Slayer", "", true)
	myAscii.Print()
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
	gameOverAscii := figure.NewColorFigure("Game over", "", "red", true)
	gameOverAscii.Print()
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
func WriteLogFile(rounds *[]RoundData) {
	execPath, err := os.Executable()

	if err != nil {
		fmt.Println(err)
	}
	execPath = filepath.Dir(execPath)
	file, err := os.Create(execPath + "/gamelog.txt")
	// Below code is only used for the time you want to use built version of the application
	//file, err := os.Create(execPath + "/gamelog.txt")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Saving a log file failed, Exiting...")
		return
	}
	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, fileErr := file.WriteString(logLine)
		if fileErr != nil {
			fmt.Println(fileErr)
			continue
		}

	}
	file.Close()
	fmt.Println("Writing into log file done.")
}
