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
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Wrong number of arguments passed, expects feed name and URL")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func handlerAddFeed(s *state, cmd command, c_user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Wrong number of arguments passed, expects feed name and URL")
	}

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    c_user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)

	feed_follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    c_user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed_follow)	

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(feeds)
	return nil
}

func handlerFollow(s *state, cmd command, c_user database.User) error {
	c_feed, err := s.DB.GetFeedWithURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	feed_follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    c_user.ID,
		FeedID:    c_feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed_follow)

	return nil
}

func handlerFollowing(s *state, cmd command, c_user database.User) error {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), c_user.ID)
	if err != nil {
		return err
	}

	fmt.Println(feeds)

	return nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		c_user, err := s.DB.GetUser(context.Background(), s.ConfigPtr.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, c_user)
	}
}

func handlerUnfollow(s *state, cmd command, c_user database.User) error {
	c_feed, err := s.DB.GetFeedWithURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	params := database.UnfollowFeedParams{
		UserID: c_user.ID,
		FeedID: c_feed.ID,
	}
	return s.DB.UnfollowFeed(context.Background(), params)
}

func scrapeFeeds(s *state) error {
	fmt.Println("Scraping")
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.DB.MarkFeedFetched(context.Background())
	if err != nil {
		return err
	}

	data, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	fmt.Println(data.Channel.Title)
	for _, item := range data.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}