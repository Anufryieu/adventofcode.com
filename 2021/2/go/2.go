package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	commands := make([]command, 0)

	for i := 0; i < len(lines); i++ {
		items := strings.Fields(lines[i])

		if len(items) == 0 {
			continue
		}

		name := items[0]
		value, err := strconv.Atoi(items[1])
		if err != nil {
			log.Fatal(err)
		}

		commands = append(commands, command{name, value})
	}

	horizontalPosition, depth := f1(commands)
	fmt.Println("1. What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(horizontalPosition * depth)

	horizontalPosition, depth = f2(commands)
	fmt.Println("2. What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(horizontalPosition * depth)
}

type command struct {
	name  string
	value int
}

func f1(commands []command) (int, int) {
	horizontalPosition := 0
	depth := 0

	for i := 0; i < len(commands); i++ {
		command := commands[i]

		switch command.name {
		case "forward":
			horizontalPosition += command.value
		case "down":
			depth += command.value
		case "up":
			depth -= command.value
		}
	}

	return horizontalPosition, depth
}

func f2(commands []command) (int, int) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for i := 0; i < len(commands); i++ {
		command := commands[i]

		switch command.name {
		case "forward":
			horizontalPosition += command.value
			depth += aim * command.value
		case "down":
			aim += command.value
		case "up":
			aim -= command.value
		}
	}

	return horizontalPosition, depth
}
