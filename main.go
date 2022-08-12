package main

import (
	"fmt"

	"github.com/ahmed-deftoner/discord-bot/pkg/bot"
	"github.com/ahmed-deftoner/discord-bot/pkg/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
