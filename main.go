package main

import (
	"8dcb/bot"
	"8dcb/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
