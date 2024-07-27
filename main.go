package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := "7499549282:AAGjfo6eVySAr_jdVTtYlW_4wOIWMl00Suo" // Replace with your bot token
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Enable debug mode
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // Ignore any non-Message updates
			continue
		}

		// Log received message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Check if the message contains the registration keyword
		if strings.HasPrefix(update.Message.Text, "Chat5 Registration Validation") {
			// Extract the username
			username := extractUsername(update.Message.Text)
			if username != "" {
				response := fmt.Sprintf("Validated username: %s", username)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No valid username found.")
				bot.Send(msg)
			}
		} else {
			// Default response
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid command. Please use 'Chat5 Registration Validation' followed by your username.")
			bot.Send(msg)
		}
	}
}

// Function to extract the username from the message text
func extractUsername(text string) string {
	// Expected format: "Chat5 Registration Validation Validate my account with username: <username>"
	parts := strings.Split(text, "username:")
	if len(parts) == 2 {
		usernameText := strings.TrimSpace(parts[1])
		username := strings.Fields(usernameText)[0]
		return username
	}
	return ""
}
