package commands

import (
	"fmt"
)

func init() {
	allCommands = append(allCommands, Command{
		Name:      "help",
		Desc:      "show command list and usage",
		AliasList: []string{},
		Handler: func(args ...string) {
			fmt.Println("Available commands are:")
			for _, command := range allCommands {
				for _, name := range command.NameList() {
					fmt.Printf(" * %s %s\n", name, command.Desc)
				}
			}
		},
		Weight: 0,
	})
}
