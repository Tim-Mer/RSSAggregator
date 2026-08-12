package main

import (
	"fmt"

	"github.com/Tim-Mer/RSSAggregator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		return
	}
	c.SetUser("tm")
	c, err = config.Read()
	if err != nil {
		return
	}
	fmt.Println(c)
}
