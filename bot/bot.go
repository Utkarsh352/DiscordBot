package bot

import (
	"8dcb/config"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func init() {
	// Retrieve the bot ID from the environment variable
	BotID = os.Getenv("BOT_ID")

	// Check if BotID is set
	if BotID == "" {
		log.Fatal("BOT_ID environment variable is not set")
	}
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

	// Keywords to check
	keywords := []string{"work", "assignment", "task", "project"}

	// Convert the message content to lowercase for case-insensitive comparison
	messageContent := strings.ToLower(m.Content)

	// Check if the message contains any of the keywords
	for _, keyword := range keywords {
		if strings.Contains(messageContent, keyword) {
			s.ChannelMessageSend(m.ChannelID, "you can contact +919306588471")
			return
		}
	}

	// Existing responses
	switch m.Content {
	case "assignment":
		s.ChannelMessageSend(m.ChannelID, "pong")
	case "hey":
		s.ChannelMessageSend(m.ChannelID, "hello")
	case "hru":
		s.ChannelMessageSend(m.ChannelID, "good u")
	case "kya kr rhe ho":
		s.ChannelMessageSend(m.ChannelID, "so rha hu")
	}
}
