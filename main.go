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

	var ops [16]byte
	for i, call := range calls {
		ops[i] = call.command.opcode<<4 + call.arg
	}

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

	/*
		for i, _ := range roms { // 0, meest rechter
			// 1, 1 naar links
			for shift, op := range ops {
				// roms[i] |= uint32(op) & (1 << uint32(shift))
				// roms[i] |= uint32(op) & uint32(i+1)
				// roms[i] |= uint32(op) & (1 << uint32(i))
				// roms[i] |= (uint32(op) & (1 << uint32(i))) << uint32(shift)

				// roms[i][-shift] = op[i]
				roms[i] |= (uint32(op) & uint32(i)) & (1 << uint32(shift))
			}
		}

		for rom, _ := range roms {
			var chars [32]byte
			for i, op := range ops {
				opstr := strconv.FormatInt(int64(op), 2)
				index := len(opstr) - rom - 1
				if index < 0 {
					index = 0
				}
				chars[31-i] = opstr[index]
			}
		}

		for _, rom := range roms {
			fmt.Printf("%032b\n", rom)
		}
	*/
}
