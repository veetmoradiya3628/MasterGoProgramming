package main

import (
	"fmt"

	"github.com/veetmoradiya3628/example/internal/color"
)

func main() {
	redText := "This is a red text"
	blueText := "This is a blue text"
	boldRedText := "This is bold red text"

	fmt.Println(color.Text(redText, color.Red))
	fmt.Println(color.Text(blueText, color.Blue))
	fmt.Println(color.Text(boldRedText, color.Red, color.Bold))

}
