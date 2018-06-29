package main

func parseInputOutputFilesFromArguments(args []string) (string, string) {
	input := args[0]
	destination := ""

	if len(args) == 2 {
		destination = args[1]
	}

	return input, destination
}

func parseArguments(args []string) {
	fileIn, fileOut = parseInputOutputFilesFromArguments(args)
}
