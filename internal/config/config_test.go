package config

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	cases := []struct {
		//key string
		val Config
	}{
		{
			//key: "https://example.com",
			val: Config{DbURL: "postgres://example"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			val, err := Read()
			if err != nil {
				t.Errorf("Got an error: %v", err)
			}
			if string(val.DbURL) != string(c.val.DbURL) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestSetUser(t *testing.T) {
	val, _ := Read()
	cases := []struct {
		key string
		val Config
	}{
		{
			key: "tm",
			val: Config{
				DbURL:           "postgres://example",
				CurrentUserName: "tm",
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			err := val.SetUser(c.key)
			if err != nil {
				t.Errorf("Got an error: %v", err)
			}
			if string(val.CurrentUserName) != string(c.val.CurrentUserName) {
				t.Errorf("expected to find value: %v\nbut got: %v", c.val.CurrentUserName, val.CurrentUserName)
				return
			}
		})
	}
}
