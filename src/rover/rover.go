package main

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Rover struct {
	x         int
	y         int
	direction Direction
}

func moveForward(rover Rover) Rover {
	switch rover.direction {
	case North:
		return Rover{x: rover.x, y: rover.y + 1, direction: North}
	case South:
		return Rover{x: rover.x, y: rover.y - 1, direction: South}
	case East:
		return Rover{x: rover.x + 1, y: rover.y, direction: East}
	case West:
		return Rover{x: rover.x - 1, y: rover.y, direction: West}
	default:
		return rover
	}
}

func moveBackward(rover Rover) Rover {
	switch rover.direction {
	case North:
		return Rover{x: rover.x, y: rover.y - 1, direction: North}
	case South:
		return Rover{x: rover.x, y: rover.y + 1, direction: South}
	case East:
		return Rover{x: rover.x - 1, y: rover.y, direction: East}
	case West:
		return Rover{x: rover.x + 1, y: rover.y, direction: West}
	default:
		return rover
	}
}

func turnLeft(rover Rover) Rover {
	switch rover.direction {
	case North:
		return Rover{x: rover.x, y: rover.y, direction: West}
	case South:
		return Rover{x: rover.x, y: rover.y, direction: East}
	case East:
		return Rover{x: rover.x, y: rover.y, direction: North}
	case West:
		return Rover{x: rover.x, y: rover.y, direction: South}
	default:
		return rover
	}
}

func turnRight(rover Rover) Rover {
	switch rover.direction {
	case North:
		return Rover{x: rover.x, y: rover.y, direction: East}
	case South:
		return Rover{x: rover.x, y: rover.y, direction: West}
	case East:
		return Rover{x: rover.x, y: rover.y, direction: South}
	case West:
		return Rover{x: rover.x, y: rover.y, direction: North}
	default:
		return rover
	}
}

func formatDirection(direction Direction) string {
	switch direction {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return ""
	}
}
