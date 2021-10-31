// Forest is a text adventure game set in a temparate forest.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetString gets reads a line from stdin. It is similar to gets in C.
func GetString() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line[:len(line)-len("\n")]
}

// GetRune gets a single rune from stdin. It is similar to getch in C.
func GetRune() rune {
	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()
	return char
}

// GetTokens reads tokens from stdin until a newline is reached
func GetTokens() []string {
   line := GetString()
   return strings.Fields(line)
}

// Prompt prints a prompt and returns input from the user.
func Prompt() string {
	fmt.Print("> ")
	return GetString()
}

// PromptTokens prints a prompt and returns an array of tokens
// input by the user. It will stop reading input when a newline
// is reached.
func PromptTokens() []string {
	fmt.Print("> ")
	return GetTokens()
}

// UnknonwCommand is used when the command input by the user is not recognized.
func UnknownCommand() {
	fmt.Println("I do not understand.")
}

// StartLocation describes the start location.
func StartLocation() {
	fmt.Println("A path winds its way from east to west through a temperate rainforest.")
	fmt.Println("The trees are huge and draped in lichen. The ground is covered by ferns")
	fmt.Println("and green moss.")
}

// UnknownLocation is used to describe the location when the player is off the grid.
// This usually means that the player is somewhere that has not been implemented yet.
func UnknownLocation() {
	fmt.Println("You are in an unknown area.")
}

// DescribeLocation described the current location. It is called when a player
// first enters a location and when the player uses the look command.
func DescribeLocation(location string) {
	if location == "start" {
		StartLocation()
	} else {
		UnknownLocation()
	}
}

// main implements the main game loop.
func main() {
	location := "start"
	last_described := ""
	for {
		if last_described != location {
			DescribeLocation(location)
			last_described = location
		}
		tokens := PromptTokens()
		var command string
		if len(tokens) > 0 {
			command = tokens[0]
		}
		switch command {
		case "quit":
			fmt.Println("Good bye")
			return
		case "character":
			fmt.Print("Enter a character: ")
			ch := GetRune()
			fmt.Printf("\nYou entered the [%c] character\n", ch)
		case "look":
			DescribeLocation(location)
		case "":
			continue
		default:
			UnknownCommand()
		}
	}
}
