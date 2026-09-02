package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Tim-Mer/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Wrong number of arguments passed, expects username only")
	}
	// Check if user is in database
	username := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User not in database")
	}
	return s.ConfigPtr.SetUser(username)
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Wrong number of arguments passed, expects username only")
	}
	_, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("Name already exists")
	}
	return s.ConfigPtr.SetUser(cmd.Args[0])
}

func handlerReset(s *state, cmd command) error {
	return s.DB.Reset(context.Background())
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get list of users")
	}
	for _, user := range users {
		line := user
		if user == s.ConfigPtr.CurrentUserName {
			line += " (current)"
		}
		fmt.Printf("- %s\n", line)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	
	return nil
}
