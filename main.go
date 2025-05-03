package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pl *ParkingLot

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		command := strings.TrimSpace(strings.ToLower(words[0]))
		if command != Add.command && pl == nil {
			fmt.Println("Parking lot not created")
			continue
		}

		switch command {
		case Add.command:
			argVal := parseArgVal(Add, words)
			if argVal != nil {
				pl = NewParkingLot(argVal[0].(int))
			}
		case Park.command:
			argVal := parseArgVal(Park, words)
			if argVal != nil {
				pl.Park(argVal[0].(string))
			}
		case Leave.command:
			argVal := parseArgVal(Leave, words)
			if argVal != nil {
				pl.Leave(argVal[0].(string), argVal[1].(int))
			}
		case Status.command:
			pl.Status()
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
