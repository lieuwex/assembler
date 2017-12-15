package main

import (
	"strconv"
	"strings"
)

type Command struct {
	name   string
	opcode byte
	hasArg bool
}

var commands = [...]Command{
	{"LDA", 4, true},
	{"LDB", 5, true},
	{"SHLA", 1, false},
	{"SHRA", 2, false},
	{"MAB", 6, false},
	{"ADDAB", 0, false},
	{"NOTA", 3, false},
	{"JC", 7, true},
}

type Call struct {
	command Command
	arg     byte
}

func makeCall(commandName string, arg byte) Call {
	var command Command
	for _, cmd := range commands {
		if cmd.name == commandName {
			command = cmd
			break
		}
	}

	return Call{command, arg}
}

func ParseLine(line string) (Call, bool, error) {
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	if len(words) == 0 {
		return Call{}, false, nil
	}

	name := words[0]
	var arg byte
	if len(words) > 1 {
		val, err := strconv.ParseInt(words[1], 0, 8)
		if err != nil {
			return Call{}, true, err
		}
		arg = byte(val)
	}
	return makeCall(name, arg), true, nil
}
