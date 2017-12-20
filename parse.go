package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Command struct {
	name   string
	opcode byte
	hasArg bool
}

var commands = [...]Command{
	{"LDA", 5, true},
	{"LDB", 6, true},
	{"SHLA", 2, false},
	{"SHRA", 3, false},
	{"MAB", 7, false},
	{"ADDAB", 1, false},
	{"NOTA", 4, false},
	{"JC", 0, true},
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

func parseNumber(str string) (byte, error) {
	base := 0
	if strings.Index(str, "0B") == 0 {
		base = 2
		str = str[2:]
	}

	res, err := strconv.ParseUint(str, base, 4)
	return byte(res), err
}

func ParseLine(line string) (Call, error) {
	words := strings.Split(line, " ")

	name := words[0]
	var arg byte
	if len(words) > 1 {
		var err error
		arg, err = parseNumber(words[1])
		if err != nil {
			return Call{}, err
		}
	}

	call, err := makeCall(name, arg)
	return call, err
}

var commentRegex = regexp.MustCompile(";.+$")

func ParseProgram(prog string) ([]Call, error) {
	var res []Call

	lines := strings.Split(prog, "\n")
	for i, line := range lines {
		line = strings.ToUpper(line)
		line = strings.TrimSpace(line)
		line = commentRegex.ReplaceAllLiteralString(line, "")

		if len(line) == 0 {
			continue
		}

		call, err := ParseLine(line)
		if err != nil {
			err = fmt.Errorf("error while parsing\n\n%d: %s\n\n%s", i, line, err.Error())
			return res, err
		}

		res = append(res, call)
	}

	return res, nil
}
