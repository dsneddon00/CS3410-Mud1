package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var allCommands = make(map[string]func(string))

// array is ordered as north, south, east, west all defaulted at 0
var globalDirections = []int{0, 0, 0, 0}


func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	initCommands()
	if err := commandLoop(); err != nil {
		log.Fatalf("%v", err)
	}

}

func commandLoop() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		doCommand(line)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("in main command loop: %v", err)
	}

	return nil
}

func addCommand(cmd string, f func(string)) {
	for i := range cmd {
		if i == 0 {
			continue
		}
		prefix := cmd[:i]
		allCommands[prefix] = f
	}
	allCommands[cmd] = f
}

func initCommands() {
	addCommand("smile", cmdSmile)
	addCommand("south", cmdSouth)
  addCommand("north", cmdNorth)
  addCommand("east", cmdEast)
  addCommand("west", cmdWest)
  addCommand("tiphat", cmdTipHat)
  addCommand("look", cmdLook)
}

func doCommand(cmd string) error {
	words := strings.Fields(cmd)
	if len(words) == 0 {
		return nil
	}

	if f, exists := allCommands[strings.ToLower(words[0])]; exists {
		f(cmd)
	} else {
		fmt.Printf("Huh?\n")
	}
	return nil
}

// split into two functions because I can't use tuples in Go
func calculateVerticalPosition(directions []int) string {
	var foo = directions[0] - directions[1]
	return strconv.Itoa(foo)
}

func calculateHorizontalPosition(directions []int) string {
	var phi = directions[2] - directions[3]
	return strconv.Itoa(phi)
}

// directional commands

func cmdNorth(s string) {
	globalDirections[0]++
  fmt.Printf("You move north. You are at position: " + calculateHorizontalPosition(globalDirections) + "," + calculateVerticalPosition(globalDirections) + " \n")
}

func cmdSouth(s string) {
	globalDirections[1]++
	fmt.Printf("You move south. You are at position: " + calculateHorizontalPosition(globalDirections) + "," + calculateVerticalPosition(globalDirections) + " \n")
}

func cmdEast(s string) {
	globalDirections[2]++
  fmt.Printf("You move east. You are at position: " + calculateHorizontalPosition(globalDirections) + "," + calculateVerticalPosition(globalDirections) + " \n")
}

func cmdWest(s string) {
	globalDirections[3]++
  fmt.Printf("You move west. You are at position: " + calculateHorizontalPosition(globalDirections) + "," + calculateVerticalPosition(globalDirections) + " \n")
}

// emote commands

func cmdSmile(s string) {
	fmt.Printf("You smile happily.\n")
}

func cmdTipHat(s string) {
  fmt.Printf("You tip your hat politely.\n")
}

// action commands

func cmdLook(s string) {
  fmt.Printf("You have looked! But there is nothing to see!\n")
}
