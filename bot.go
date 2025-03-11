package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {
	// load the bot token from environment variables
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("No token found! Set DISCORD_BOT_TOKEN environment variable.")
	}

	// creates a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	// add message handler
	dg.AddHandler(messageCreate)

	// open a WebSocket connection to Discord
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	fmt.Println("Bot is now running! Press CTRL+C to exit.")

	// keep bot running
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore bot's own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println("Received message:", m.Content) // debugging log

	// check if the message starts with "!bet"
	if strings.HasPrefix(m.Content, "!bet") {
		parts := strings.Fields(m.Content) // splits message by spaces

		// validate command format
		if len(parts) != 4 {
			s.ChannelMessageSend(m.ChannelID, "Invalid format! Use: `!bet <name> <amount> <over/under/7>`")
			return
		}

		player := parts[1]
		betAmount, err := strconv.Atoi(parts[2]) // convert amount to integer
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Invalid bet amount! Enter a number.")
			return
		}

		betType := strings.ToLower(parts[3]) // get "over", "under", or "7"

		// call Roll function from game.go
		result := Roll(betType, betAmount)

		// send the result message to Discord
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s placed a bet!\n%s", player, result))
	}
}
