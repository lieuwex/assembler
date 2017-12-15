package main

import (
	"fmt"
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

func makeCall(commandName string, arg byte) (Call, error) {
	var command Command
	has := false

	for _, cmd := range commands {
		if cmd.name == commandName {
			command = cmd
			has = true
			break
		}
	}

	var err error = nil
	if !has {
		err = fmt.Errorf("no command '%s' found", commandName)
	}
	return Call{command, arg}, err
}

func ParseLine(line string) (Call, error) {
	words := strings.Split(line, " ")

	name := words[0]
	var arg byte
	if len(words) > 1 {
		val, err := strconv.ParseInt(words[1], 0, 8)
		if err != nil {
			return Call{}, err
		}
		arg = byte(val)
	}

	call, err := makeCall(name, arg)
	return call, err
}

func ParseProgram(prog string) ([]Call, error) {
	var res []Call

	lines := strings.Split(prog, "\n")
	for _, line := range lines {
		line = strings.ToUpper(line)
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		call, err := ParseLine(line)
		if err != nil {
			return res, err
		}

		res = append(res, call)
	}

	return res, nil
}
