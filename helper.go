package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseArgVal(command Command, words []string) (result []any) {
	if len(words)-1 != len(command.inputs) {
		fmt.Printf("Invalid %v command: %v\n", command.command, words[0])
		return nil
	}

	for i, val := range words[1:] {
		val = strings.TrimSpace(val)
		if command.inputs[i].isChar {
			result = append(result, val)
		} else {
			parsedVal, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Invalid %v: %v\n", command.inputs[i].label, val)
				return nil
			}
			result = append(result, parsedVal)
		}
	}

	return
}
