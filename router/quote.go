package router

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/valentinlegal/chef-michel-dumas-bot/data"
)

func (m *Mux) Quote(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	log.Println("[INFO] Quote router called")

	quote, err := pickQuote()
	if err != nil {
		log.Fatalf("[ERROR] Could not pick a quote: %v", err)
	}

	_, err = ds.ChannelMessageSend(dm.ChannelID, quote)
	if err != nil {
		log.Fatalf("[ERROR] Could not send Discord message: %v", err)
	}
}

func pickQuote() (string, error) {
	quotes := data.Quotes()
	if len(quotes) == 0 {
		return "", fmt.Errorf("array must contain at least one quote. got=%v", len(quotes))
	}

	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(quotes))

	return quotes[i], nil
}
