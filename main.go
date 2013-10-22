package main

import (
	"fmt"
	"github.com/robmerrell/comandante"
	"github.com/robmerrell/comandante_example/cmds"
	"os"
)

func main() {
	bin := comandante.New("comandante_example", "Example program showing how to use Comandante")
	bin.IncludeHelp() // autogenerate a "help" command

	// list command
	listCmd := comandante.NewCommand("list", "list files in a directory", cmds.ListAction)
	listCmd.Documentation = cmds.ListDoc
	bin.RegisterCommand(listCmd)

	// http get command
	getCmd := comandante.NewCommand("get", "GETs the contents of a webpage", cmds.GetAction)
	getCmd.FlagInit = cmds.GetFlagHandler // use the flag package to get a url
	getCmd.Documentation = cmds.GetDoc
	bin.RegisterCommand(getCmd)

	// run the commands
	if err := bin.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
