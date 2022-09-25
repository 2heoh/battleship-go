package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"battleship-go/cmd/console"
	"battleship-go/cmd/gamecontroller"
	"battleship-go/cmd/contracts"
	"battleship-go/cmd/telemetryclient"
)

var (
	enemyFleet []*contracts.Ship
	myFleet    []*contracts.Ship
	scanner    *bufio.Scanner
	apptelemetryclient *telemetryclient.TelemetryClient
)
var printer = console.ColoredPrinter(1, false).Background(contracts.BLACK).Foreground(contracts.WHITE).Build()

func main() {
	apptelemetryclient = telemetryclient.NewTelemetryClient()
	apptelemetryclient.TrackEvent("ApplicationStarted", map[string]string{"Technology": "GOLang"})
	scanner = bufio.NewScanner(os.Stdin)

	printer.SetForegroundColor(contracts.MAGENTA)
	printer.Println("                                     |__")
	printer.Println("                                     |\\/")
	printer.Println("                                     ---")
	printer.Println("                                     / | [")
	printer.Println("                              !      | |||")
	printer.Println("                            _/|     _/|-++'")
	printer.Println("                        +  +--|    |--|--|_ |-")
	printer.Println("                     { /|__|  |/\\__|  |--- |||__/")
	printer.Println("                    +---------------___[}-_===_.'____                 /\\")
	printer.Println("                ____`-' ||___-{]_| _[}-  |     |_[___\\==--            \\/   _")
	printer.Println(" __..._____--==/___]_|__|_____________________________[___\\==--____,------' .7")
	printer.Println("|                        Welcome to Battleship                         BB-61/")
	printer.Println(" \\_________________________________________________________________________|")
	printer.Println("")
	printer.SetForegroundColor(contracts.WHITE)

	initializeGame()

	startGame()
}

func startGame() {
	printer.Print("\033[2J\033[;H")
	printer.Println("                  __")
	printer.Println("                 /  \\")
	printer.Println("           .-.  |    |")
	printer.Println("   *    _.-'  \\  \\__/")
	printer.Println("    \\.-'       \\")
	printer.Println("   /          _/")
	printer.Println("  |      _  /\" \"")
	printer.Println("  |     /_\\'")
	printer.Println("   \\    \\_/")
	printer.Println("    \" \"\" \"\" \"\" \"")

	for {
		printer.Println("")
		printer.Println("Player, it's your turn")
		printer.Println("Enter coordinates for your shot :")
		var isHit bool
		for scanner.Scan() {
			input := scanner.Text()
			if input != "" {
				position := parsePosition(input)
				var err error
				isHit, err = gamecontroller.CheckIsHit(enemyFleet, position)
				apptelemetryclient.TrackEvent("Player_ShootPosition", map[string]string{"Position": input, "IsHit": strconv.FormatBool(isHit) });
				if err != nil {
					printer.Printf("Error: %s\n", err)
				}
				break
			}
		}

		if isHit {
			beep()
			printer.Println("                \\         .  ./")
			printer.Println("              \\      .:\" \";'.:..\" \"   /")
			printer.Println("                  (M^^.^~~:.'\" \").")
			printer.Println("            -   (/  .    . . \\ \\)  -")
			printer.Println("               ((| :. ~ ^  :. .|))")
			printer.Println("            -   (\\- |  \\ /  |  /)  -")
			printer.Println("                 -\\  \\     /  /-")
			printer.Println("                   \\  \\   /  /")
		}

		if isHit {
			printer.Println("Yeah ! Nice hit !")
		} else {
			printer.Println("Miss")
		}

		position := getRandomPosition()
		var err error
		isHit, err = gamecontroller.CheckIsHit(myFleet, position)
		apptelemetryclient.TrackEvent("Computer_ShootPosition", map[string]string{"Position": tostring(*position), "IsHit": strconv.FormatBool(isHit) });
		
		if err != nil {
			printer.Printf("Error: %s\n", err)
		}
		printer.Println("")

		result := "hit your ship !"
		if !isHit {
			result = "miss"
		}
		printer.Printf("Computer shoot in %s%d and %s\n", position.Column, position.Row, result)
		if isHit {
			beep()
			printer.Println("                \\         .  ./")
			printer.Println("              \\      .:\" \";'.:..\" \"   /")
			printer.Println("                  (M^^.^~~:.'\" \").")
			printer.Println("            -   (/  .    . . \\ \\)  -")
			printer.Println("               ((| :. ~ ^  :. .|))")
			printer.Println("            -   (\\- |  \\ /  |  /)  -")
			printer.Println("                 -\\  \\     /  /-")
			printer.Println("                   \\  \\   /  /")
		}
	}
}

func parsePosition(input string) *contracts.Position {
	ltr := strings.ToUpper(string(input[0]))
	number, _ := strconv.Atoi(string(input[1]))
	return &contracts.Position{
		Column: contracts.FromString(ltr),
		Row:    number,
	}
}

func tostring(p contracts.Position) string {
	return p.Column.String() + strconv.FormatInt(int64(p.Row), 10)
}

func beep() {
	fmt.Print("\007")
}

func initializeGame() {
	initializeMyFleet()

	initializeEnemyFleet()
}

func initializeMyFleet() {
	//reader := bufio.NewReader(os.Stdin)
	//scanner := bufio.NewScanner(os.Stdin)
	myFleet = gamecontroller.InitializeShips()

	printer.Println("Please position your fleet (Game board has size from A to H and 1 to 8) :")

	for _, ship := range myFleet {
		printer.Println("")
		printer.Printf("Please enter the positions for the %s (size: %d)", ship.Name, ship.Size)
		printer.Println("")

		for i := 1; i <= ship.Size; i++ {
			printer.Printf("Enter position %d of %d (i.e A3):\n", i, ship.Size)

			for scanner.Scan() {
				positionInput := scanner.Text()
				if positionInput != "" {
					ship.AddPosition(positionInput)
					apptelemetryclient.TrackEvent("Player_PlaceShipPosition", map[string]string{"Position": positionInput, "Ship": ship.Name, "PositionInShip": strconv.FormatInt(int64(len(ship.Positions)), 10) });
					break
				}
			}
		}
	}
}

func getRandomPosition() *contracts.Position {
	rows := 8
	lines := 8
	letter := contracts.Letter(rand.Intn(lines-1) + 1)
	number := rand.Intn(rows-1) + 1
	position := &contracts.Position{Column: letter, Row: number}
	return position
}

func initializeEnemyFleet() {
	enemyFleet = gamecontroller.InitializeShips()

	enemyFleet[0].SetPositions(&contracts.Position{Column: contracts.B, Row: 4})
	enemyFleet[0].SetPositions(&contracts.Position{Column: contracts.B, Row: 5})
	enemyFleet[0].SetPositions(&contracts.Position{Column: contracts.B, Row: 6})
	enemyFleet[0].SetPositions(&contracts.Position{Column: contracts.B, Row: 7})
	enemyFleet[0].SetPositions(&contracts.Position{Column: contracts.B, Row: 8})

	enemyFleet[1].SetPositions(&contracts.Position{Column: contracts.E, Row: 6})
	enemyFleet[1].SetPositions(&contracts.Position{Column: contracts.E, Row: 7})
	enemyFleet[1].SetPositions(&contracts.Position{Column: contracts.E, Row: 8})
	enemyFleet[1].SetPositions(&contracts.Position{Column: contracts.E, Row: 9})

	enemyFleet[2].SetPositions(&contracts.Position{Column: contracts.A, Row: 3})
	enemyFleet[2].SetPositions(&contracts.Position{Column: contracts.B, Row: 3})
	enemyFleet[2].SetPositions(&contracts.Position{Column: contracts.C, Row: 3})

	enemyFleet[3].SetPositions(&contracts.Position{Column: contracts.F, Row: 8})
	enemyFleet[3].SetPositions(&contracts.Position{Column: contracts.G, Row: 8})
	enemyFleet[3].SetPositions(&contracts.Position{Column: contracts.H, Row: 8})

	enemyFleet[4].SetPositions(&contracts.Position{Column: contracts.C, Row: 5})
	enemyFleet[4].SetPositions(&contracts.Position{Column: contracts.C, Row: 6})
}
