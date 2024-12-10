package main

import (
	"flag"
	"fmt"
	"linuxcommands/basic"
	"linuxcommands/help"
)

type Commands interface {
	GetCommands() map[string]string
}

func main() {
	command := flag.String("c", "", "Command to execute")
	flag.StringVar(command, "command", "", "Command to execute")
	flag.Parse()

	var c Commands
	switch *command {
	case basic.BasicName:
		c = basic.NewBasic()
	case help.HelpName:
		c = help.NewHelp()
	default:
		c = help.NewHelp()
	}

	commands := c.GetCommands()
	for k, v := range commands {
		fmt.Println(fmt.Sprintf("With %s", k), fmt.Sprintf("-- you can %s", v))
	}
}
