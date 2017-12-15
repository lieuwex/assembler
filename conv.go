package main

func CallsToOps(calls []Call) []byte {
	res := make([]byte, len(calls))

	for i, call := range calls {
		res[i] = call.command.opcode<<4 + call.arg
	}

	return res
}
