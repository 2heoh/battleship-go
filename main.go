package main

import (
	"fmt"

	"gitlab.com/2heoh/go-battleship/console"
)

func main() {
	console.SetForegroundColor(console.MAGENTA)
	fmt.Println("Go battleship!")
	console.ResetForegroundColor()
	fmt.Println("Go battleship!")
}

