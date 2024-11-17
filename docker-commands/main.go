package main

import (
	"docker-commands/containers"
	"docker-commands/general"
	"docker-commands/help"
	"docker-commands/hub"
	"docker-commands/images"
	"flag"
	"fmt"
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
	case images.ImagesName:
		c = images.NewImages()
	case help.HelperName:
		c = help.NewHelper()
	case hub.HubName:
		c = hub.NewHub()
	case general.GeneralName:
		c = general.NewGeneral()
	case containers.ContainersName:
		c = containers.NewContainers()
	default:
		return
	}

	commands := c.GetCommands()
	for k, v := range commands {
		fmt.Println(fmt.Sprintf("With %s", k), fmt.Sprintf("-- you can %s", v))
	}
}
