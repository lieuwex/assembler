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
		panic(err)
	} else if len(calls) > 16 {
		err = fmt.Errorf("program too long (got %d, max is %d)", len(calls), 16)
		panic(err)
	}

	ops := CallsToOps(calls)

	var roms [7]uint32
	for shift, op := range ops {
		romCol := 1 << uint32(shift)
		for rom, _ := range roms {
			opCol := 1 << uint32(rom)
			if op&byte(opCol) != 0 {
				roms[rom] |= uint32(romCol)
			}
		}
	}

	for _, rom := range roms {
		fmt.Printf("%08X\n", rom)
	}
}
