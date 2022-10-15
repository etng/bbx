package main

import (
	"fmt"
	"github.com/etng/bbx/commands"
	"github.com/etng/bbx/helpers"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	if !helpers.GetBoolEnv("BBX_NO_BANNER") {
		fmt.Println("Welcome to use BBX")
	}

	cmdName := "help"
	offset := 1
	if len(os.Args) > 1 {
		cmdName = os.Args[1]
		offset = 2
	}
	commands.Call(cmdName, os.Args[offset:]...)
}
