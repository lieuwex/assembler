package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	calls, err := ParseProgram(string(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	} else if len(calls) > 16 {
		err = fmt.Errorf("program too long (got %d, max is %d)", len(calls), 16)
		panic(err)
	}

	ops := CallsToOps(calls)
	roms := BuildRom(ops)

	for _, rom := range roms {
		fmt.Printf("%08X\n", rom)
	}
}
