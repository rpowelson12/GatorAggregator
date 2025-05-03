package main

import (
	"log"
	"os"

	"github.com/rpowelson12/GatorAggregator/internal/config"
)

func main() {
	// 1. Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	cfgState := state{&cfg}
	commandList := newCommands()

	commandList.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("error with input")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := command{
		name: commandName,
		args: commandArgs,
	}

	err = commandList.run(&cfgState, cmd)
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}
