package main

import (
	"log"
	"os"

	"github.com/johnbellone/skel/command"
	"github.com/mitchellh/cli"
)

func main() {
	log.SetOutput(os.Stderr)
	c := cli.NewCLI("skel", Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"new": func() (cli.Command, error) {
			return &command.NewCommand{}, nil
		},
		"sync": func() (cli.Command, error) {
			return &command.SyncCommand{}, nil
		},
		"up": func() (cli.Command, error) {
			return &command.UpCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
