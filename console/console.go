package console

import "fmt"



func SetForegroundColor(color string)  {
	fmt.Print(color)
}

func ResetForegroundColor() {
	fmt.Print(DEFAULT_GREY)
}
