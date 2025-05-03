package main

import (
	"errors"
	"fmt"
	"github.com/rpowelson12/GatorAggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("you must give a name")
	}
	s.config.SetUser(cmd.args[0])
	fmt.Println("User has been set")
	return nil
}
