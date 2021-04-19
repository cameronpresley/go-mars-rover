package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r := Rover{x: 0, y: 0, direction: North}
	input := ""
	for ok := true; ok; ok = (input != "q") {
		input, _ = getInput(reader, "Choose a command Move (f)orward, Move (b)ackward, Turn (l)eft, Turn (r)ight, or (q)uit?")
		input = strings.ToLower(input)
		switch input {
		case "f":
			r = moveForward(r)
		case "b":
			r = moveBackward(r)
		case "l":
			r = turnLeft(r)
		case "r":
			r = turnRight(r)
		case "q":
		default:
		}
		fmt.Printf("Rover is at (%d,%d) facing %s\n", r.x, r.y, formatDirection(r.direction))
	}
}

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	s, e := reader.ReadString('\n')
	if e != nil {
		return "", e
	}
	return strings.TrimSpace(s), nil
}
