package bot

import (
	"8dcb/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func init() {
	BotID = "791960000413941760"
}

func Start() {
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal(err)
	}

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if m.Content == "hey" {
		s.ChannelMessageSend(m.ChannelID, "hello")
	}
	if m.Content == "hru" {
		s.ChannelMessageSend(m.ChannelID, "good u")
	}
}
