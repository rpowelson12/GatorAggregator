package main

import "errors"

type commands struct {
	commands map[string]func(*state, command) error
}

func newCommands() *commands {
	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	return &commands
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("command not registered")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
