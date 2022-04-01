package router

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/valentinlegal/chef-michel-dumas-bot/data"
)

var Embed *discordgo.MessageEmbed

func (m *Mux) Help(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	log.Println("[INFO] Help router called")

	// If embed already set, return it
	if Embed != nil {
		_, err := ds.ChannelMessageSendEmbed(dm.ChannelID, Embed)
		if err != nil {
			log.Fatalf("[ERROR] Could not send Discord message: %v", err)
		}
		return
	}

	avatarUrl := os.Getenv("BOT_AVATAR_URL")
	title := os.Getenv("BOT_TITLE")
	url := os.Getenv("BOT_URL")
	version := os.Getenv("BOT_VERSION")

	// Lists the main bot commands
	fields, err := CommandFields()
	if err != nil {
		log.Fatalf("[ERROR] Could not generate commands fields: %v", err)
	}

	Embed = &discordgo.MessageEmbed{
		Title:       title,
		URL:         url,
		Description: "**Toutes les commandes du bot :**",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatarUrl,
		},
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("%s - Version %s", title, version),
			IconURL: avatarUrl,
		},
	}

	_, err = ds.ChannelMessageSendEmbed(dm.ChannelID, Embed)
	if err != nil {
		log.Fatalf("[ERROR] Could not send Discord message: %v", err)
	}
}

// CommandFields generates command fields based on the data file.
// It return an array of Discord embed field with the commands and the associated description.
func CommandFields() ([]*discordgo.MessageEmbedField, error) {
	var fields []*discordgo.MessageEmbedField

	commands := data.Commands()
	if len(commands) == 0 {
		return nil, fmt.Errorf("array must contain at least one command. got=%v", len(commands))
	}

	for _, c := range commands {
		var name string

		for i, k := range c.Keys {
			name += "`!chef"

			if k != "" {
				name += " " + k
			}

			name += "`"

			if i != len(c.Keys)-1 {
				name += " / "
			}
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  name,
			Value: c.Description,
		})
	}

	return fields, nil
}
