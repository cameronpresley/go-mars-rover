package main

import (
	"testing"
)

type Location struct {
	x int
	y int
}

func TestRoverMovesForward(t *testing.T) {
	directions := [4]Direction{North, South, East, West}
	locations := [4]Location{Location{x: 0, y: 1}, Location{x: 0, y: -1}, Location{x: 1, y: 0}, Location{x: -1, y: 0}}

	for i := 0; i < 4; i++ {
		r := moveForward(Rover{x: 0, y: 0, direction: directions[i]})
		if r.direction != directions[i] {
			t.Errorf("Rover expected to face %s, but was facing %s", formatDirection(directions[i]), formatDirection(r.direction))
		}
		expectedLocation := locations[i]
		if r.x != expectedLocation.x {
			t.Errorf("Rover expected to have X of %d, but had %d", expectedLocation.x, r.x)
		}
		if r.y != expectedLocation.y {
			t.Errorf("Rover expected to have Y of %d, but had %d", expectedLocation.y, r.y)
		}
	}
}

func TestRoverMovesBackward(t *testing.T) {
	directions := [4]Direction{North, South, East, West}
	locations := [4]Location{Location{x: 0, y: -1}, Location{x: 0, y: 1}, Location{x: -1, y: 0}, Location{x: 1, y: 0}}

	for i := 0; i < 4; i++ {
		r := moveBackward(Rover{x: 0, y: 0, direction: directions[i]})
		if r.direction != directions[i] {
			t.Errorf("Rover expected to face %s, but was facing %s", formatDirection(directions[i]), formatDirection(r.direction))
		}
		expectedLocation := locations[i]
		if r.x != expectedLocation.x {
			t.Errorf("Rover expected to have X of %d, but had %d", expectedLocation.x, r.x)
		}
		if r.y != expectedLocation.y {
			t.Errorf("Rover expected to have Y of %d, but had %d", expectedLocation.y, r.y)
		}
	}
}

func TestRoverTurnsLeft(t *testing.T) {
	directions := [4]Direction{North, South, East, West}
	expectedDirections := [4]Direction{West, East, North, South}
	for i := 0; i < 4; i++ {
		r := turnLeft(Rover{x: 0, y: 0, direction: directions[i]})
		if r.direction != expectedDirections[i] {
			t.Errorf("Rover expected to face %s, but was facing %s", formatDirection(directions[i]), formatDirection(r.direction))
		}
		if r.x != 0 {
			t.Errorf("Rover expected to have X of 0, but had %d", r.x)
		}
		if r.y != 0 {
			t.Errorf("Rover expected to have Y of 0, but had %d", r.y)
		}
	}
}

func TestRoverTurnsRight(t *testing.T) {
	directions := [4]Direction{North, South, East, West}
	expectedDirections := [4]Direction{East, West, South, North}
	for i := 0; i < 4; i++ {
		r := turnRight(Rover{x: 0, y: 0, direction: directions[i]})
		if r.direction != expectedDirections[i] {
			t.Errorf("Rover expected to face %s, but was facing %s", formatDirection(directions[i]), formatDirection(r.direction))
		}
		if r.x != 0 {
			t.Errorf("Rover expected to have X of 0, but had %d", r.x)
		}
		if r.y != 0 {
			t.Errorf("Rover expected to have Y of 0, but had %d", r.y)
		}
	}
}

func TestFormatDirection(t *testing.T) {
	directions := [4]Direction{North, South, East, West}
	expectedMessage := [4]string{"North", "South", "East", "West"}

	for i := 0; i < 4; i++ {
		actualMessage := formatDirection(directions[i])
		if expectedMessage[i] != actualMessage {
			t.Errorf("Expected %s, but got %s instead", expectedMessage[i], actualMessage)
		}
	}
}
