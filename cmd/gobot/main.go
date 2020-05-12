package main

import (
	"fmt"

	"github.com/cyberpunkprogrammer/gobot/pkg/bot"
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

func main() {

	err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
}
