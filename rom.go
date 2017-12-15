package main

func BuildRom(ops []byte) (roms [7]uint32) {
	for shift, op := range ops {
		romCol := 1 << uint32(shift)
		for rom, _ := range roms {
			opCol := 1 << uint32(rom)
			if op&byte(opCol) != 0 {
				roms[rom] |= uint32(romCol)
			}
		}
	}

	return roms
}
