package config

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	cases := []struct {
		val Config
	}{
		{
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
			if string(c.key) != string(c.val.CurrentUserName) {
				t.Errorf("expected to find value: %v\nbut got: %v", c.key, c.val.CurrentUserName)
				return
			}
		})
	}
}

func TestHandlerLogin(t *testing.T) {

	config, _ := Read()
	configPtr := &config
	st := state{}
	st.configPtr = configPtr

	cmd := command{}
	cmd.name = "login"

	cases := []struct {
		key string
		val *state
	}{
		{
			key: "test",
			val: &st,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cmd.args = []string{c.key}
			err := handlerLogin(c.val, cmd)

			if err != nil {
				t.Errorf("Got an error: %v", err)
			}

			if string(c.val.configPtr.CurrentUserName) != string(c.key) {
				t.Errorf("expected to find value: %v\nbut got: %v", c.key, c.val.configPtr.CurrentUserName)
				return
			}
		})
	}
}
