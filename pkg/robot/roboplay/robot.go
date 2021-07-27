package roboplay

import (
	"fmt"

	"github.com/bvinayb/play-robo/pkg/robot/direction"
	"github.com/bvinayb/play-robo/pkg/robot/table"
)

type Robot struct {
	X, Y      int
	Direction string
}

func Place(x int, y int, direction string) (*Robot, error) {
	if !table.IsValidPosition(x, y) {
		return nil, fmt.Errorf("not a valid position")
	}
	return &Robot{x, y, direction}, nil
}

func Move(robot *Robot) *Robot {
	x, y, direction := robot.X, robot.Y, robot.Direction

	switch direction {
	case "NORTH":
		y += 1
	case "EAST":
		x += 1
	case "SOUTH":
		y -= 1
	case "WEST":
		x -= 1
	}
	if !table.IsValidPosition(x, y) {
		return robot
	}
	return &Robot{x, y, direction}
}

func Left(robot *Robot) *Robot {
	return &Robot{robot.X, robot.Y, direction.Left(robot.Direction)}
}

func Right(robot *Robot) *Robot {
	return &Robot{robot.X, robot.Y, direction.Right(robot.Direction)}
}

func Report(robot *Robot) string {
	return fmt.Sprintf("%d,%d,%s", robot.X, robot.Y, robot.Direction)
}
