package router

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/valentinlegal/chef-michel-dumas-bot/data"
)

func (m *Mux) Gif(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	log.Println("[INFO] Gif router called")

	gif, err := pickGif()
	if err != nil {
		log.Fatalf("[ERROR] Could not pick a gif: %v", err)
	}

	_, err = ds.ChannelMessageSend(dm.ChannelID, gif)
	if err != nil {
		log.Fatalf("[ERROR] Could not send Discord message: %v", err)
	}
}

func pickGif() (string, error) {
	gifs := data.Gifs()
	if len(gifs) == 0 {
		return "", fmt.Errorf("array must contain at least one GIF. got=%v", len(gifs))
	}

	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(gifs))

	return gifs[i], nil
}
