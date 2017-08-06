package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func whatPace(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf(
			"Wrong number of arguments - %v. 2 expected", len(args),
		)
	}

	km, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf(
			"Wrong value for km - %v. Integer expected", args[0],
		)
	}
	time := args[1]

	parsedTime := strings.Split(time, ":")
	parsedTimeLen := len(parsedTime)

	var hours, minutes int

	if parsedTimeLen == 2 {
		hours, err = strconv.Atoi(parsedTime[0])
		if err != nil {
			return "", fmt.Errorf(
				"Wrong value for hours - %v. Integer expected", parsedTime[0],
			)
		}
		minutes, err = strconv.Atoi(parsedTime[1])
		if err != nil {
			return "", fmt.Errorf(
				"Wrong value for minutes - %v. Integer expected",
				parsedTime[1],
			)
		}
	} else if parsedTimeLen == 1 {
		hours = 0
		minutes, err = strconv.Atoi(parsedTime[0])
		if err != nil {
			return "", fmt.Errorf(
				"Wrong value for minutes - %v. Integer expected",
				parsedTime[0],
			)
		}
	} else {
		return "", fmt.Errorf("Wrong time format. H:M or M expected")
	}

	seconds := (hours*3600 + minutes*60) / km
	paceMinutes := seconds / 60
	paceSeconds := seconds % 60

	var pace string
	if paceMinutes > 0 {
		pace = fmt.Sprintf("%vm %0.2vs", paceMinutes, paceSeconds)
	} else {
		pace = fmt.Sprintf("%0.2v", paceSeconds)
	}

	pace = fmt.Sprintf(
		"Pace to finish %v km in %s - %s per km", km, time, pace,
	)
	return pace, nil
}

func main() {
	pace, err := whatPace(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(pace)
}
