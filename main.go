package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Tim-Mer/RSSAggregator/internal/config"
	"github.com/Tim-Mer/RSSAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	ConfigPtr *config.Config
	DB        *database.Queries
}

func main() {
	args := os.Args
	cmd := command{
		Name: os.Args[1],
		Args: args[2:],
	}

	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state := state{ConfigPtr: &c}
	commands := commands{}
	if err := commands.Initialise(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", state.ConfigPtr.DbURL)
	dbQueries := database.New(db)
	state.DB = dbQueries

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

//"postgres://postgres:postgres@localhost:5432/gator"
