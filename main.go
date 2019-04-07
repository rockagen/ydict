package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	proxy string
)

func main() {
	//Check & load .env file
	loadEnv()

	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		displayUsage()
		os.Exit(0)
	}

	words, withVoice, withMore, isQuiet, interaction := parseArgs(os.Args[1:])

	if interaction {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			Query(strings.Fields(text),withMore,withMore,isQuiet,len(words) > 1)
		}

	}else{
		Query(words, withVoice, withMore, isQuiet, len(words) > 1)
	}
}
