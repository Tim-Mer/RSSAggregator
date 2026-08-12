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
			val: Config{},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			val := Read()
			if string(val.DbURL) != string(c.val.DbURL) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
