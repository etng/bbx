package commands

import (
	"sort"
	"sync"
)

type Command struct {
	Name      string
	Desc      string
	AliasList []string
	Handler   func(args ...string)
	Weight    int
}

func (c Command) NameList() []string {
	list := append([]string{c.Name}, c.AliasList...)
	sort.Strings(list)
	return list
}

var allCommands []Command
var indexOnce sync.Once
var indexedCommands map[string]Command

func Index() {
	indexOnce.Do(func() {
		sort.Slice(allCommands, func(i, j int) bool {
			if allCommands[i].Weight == allCommands[j].Weight {
				return allCommands[i].Name < allCommands[j].Name
			}
			return allCommands[i].Weight < allCommands[j].Weight
		})
		indexedCommands = make(map[string]Command, len(allCommands))
		for _, command := range allCommands {
			for _, name := range command.NameList() {
				indexedCommands[name] = command
			}
		}
	})
}

func Call(name string, args ...string) {
	Index()
	cmd, ok := indexedCommands[name]
	if !ok {
		indexedCommands["help"].Handler(args...)
	} else {
		cmd.Handler(args...)
	}
}
