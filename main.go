package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bvinayb/play-robo/pkg/robot/direction"
	"github.com/bvinayb/play-robo/pkg/robot/roboplay"
)

func main() {
	fmt.Println("Playing Toy Robot....")
	var live_robot *roboplay.Robot
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		commandWithLineReturn, _ := reader.ReadString('\n')
		command := strings.TrimSpace(commandWithLineReturn)
		command = strings.ToUpper(command)
		live_robot, err = runCommand(command, live_robot)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func runCommand(command string, my_robot *roboplay.Robot) (*roboplay.Robot, error) {
	if regexp.MustCompile("PLACE [0-9]+,[0-9]+,[A-Z]+").MatchString(command) {
		return runPlaceCommand(command, my_robot)
	}
	if my_robot == nil {
		return nil, fmt.Errorf("Please place the robot, its not in place")
	}
	switch {
	case command == "MOVE":
		return roboplay.Move(my_robot), nil
	case command == "LEFT":
		return roboplay.Left(my_robot), nil
	case command == "RIGHT":
		return roboplay.Right(my_robot), nil
	case command == "REPORT":
		fmt.Println(roboplay.Report(my_robot))
		return my_robot, nil
	case command == "Q":
		os.Exit(0)
	case command == "" || command == "\n":
		return my_robot, nil
	}
	return my_robot, fmt.Errorf("Command not recognized: %s", command)
}

func runPlaceCommand(command string, my_robot *roboplay.Robot) (*roboplay.Robot, error) {
	args_string := strings.Fields(command)[1]
	args := strings.Split(args_string, ",")
	x, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	placing_direction := args[2]
	if !direction.IsValidDirection(placing_direction) {
		return my_robot, fmt.Errorf("Invalid Direction")
	}
	new_robot, err := roboplay.Place(x, y, placing_direction)
	if err != nil {
		return my_robot, fmt.Errorf(err.Error())
	}
	return new_robot, nil
}
