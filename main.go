package main

import (
	"fmt"
	"github.com/rpowelson12/GatorAggregator/internal/config"
	"log"
)

func main() {
	// 1. Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// 2. Set the current user
	err = cfg.SetUser("Randy")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	// 3. Read the config file again
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error updating config: %v", err)
	}
	// 4. Print the contents
	fmt.Printf("%+v", updatedCfg)
}
