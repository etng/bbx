package commands

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"github.com/etng/bbx/helpers"
	"io/ioutil"
	"os"
)

func init() {
	allCommands = append(allCommands, Command{
		Name:      "markdown",
		Desc:      "render markdown in terminal",
		AliasList: []string{"md"},
		Handler: func(args ...string) {
			if len(args) < 1 {
				fmt.Println("you need pass filename")
				os.Exit(1)
			}
			b, e := ioutil.ReadFile(args[0])
			helpers.PanicIf(e)
			out, err := glamour.Render(string(b), "dark")
			helpers.PanicIf(err)
			fmt.Print(out)
		},
		Weight: 0,
	})
}
