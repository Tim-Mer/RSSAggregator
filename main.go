package main

import (
	"fmt"
	"os"

	"github.com/Tim-Mer/RSSAggregator/internal/config"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		fmt.Println("Error: Not enough args passed")
		os.Exit(1)
	}
	cmd := config.Command{
		Name: os.Args[1],
		Args: args[2:],
	}

	c, err := config.Read()
	if err != nil {
		os.Exit(1)
	}

	state := config.State{ConfigPtr: &c}
	commands := config.Commands{}
	if err := commands.Initialise(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = commands.Run(&state, cmd)

	os.Exit(0)
}
