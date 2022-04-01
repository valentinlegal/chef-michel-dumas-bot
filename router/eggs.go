package router

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Eggs(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	log.Println("[INFO] Eggs router called")

	var msg string

	switch ctx.Content {
	case "minecraft", "mc", "fanta":
		msg = "https://youtu.be/2P1PHiFBQT0"
	case "remix", "dj":
		msg = "https://youtu.be/O1a7RhkMxPU"
	}

	_, err := ds.ChannelMessageSend(dm.ChannelID, msg)
	if err != nil {
		log.Fatalf("[ERROR] Could not send Discord message: %v", err)
	}
}
