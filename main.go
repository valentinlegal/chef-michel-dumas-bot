package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"github.com/valentinlegal/chef-michel-dumas-bot/data"
	"github.com/valentinlegal/chef-michel-dumas-bot/router"
)

var Token string
var Router = router.New()

func init() {
	// Loads .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[ERROR] Could not load .env file: %v", err)
	}

	// Assigns bot's token
	Token = os.Getenv("BOT_TOKEN")

	// ===========
	// R O U T E S
	// ===========

	// Sends a quote
	Router.Route("quote", "Send a quote.", Router.Quote)
	Router.Route("michel", "Send a quote.", Router.Quote)

	// Sends a GIF
	Router.Route("gif", "Send a GIF.", Router.Gif)
	Router.Route("supernickel", "Send a GIF.", Router.Gif)
	Router.Route("super", "Send a GIF.", Router.Gif)
	Router.Route("nickel", "Send a GIF.", Router.Gif)

	// Sends main available commands and information
	Router.Route("help", "Send main available commands and information.", Router.Help)

	// Some easter eggs
	Router.Route("minecraft", `Send the song "Minecraft - Antoine Daniel".`, Router.Eggs)
	Router.Route("mc", `Send the song "Minecraft - Antoine Daniel".`, Router.Eggs)
	Router.Route("fanta", `Send the song "Minecraft - Antoine Daniel".`, Router.Eggs)
	Router.Route("remix", `Send the song "Goudja On The Floor - Risitas ft. Cocinero Dumas".`, Router.Eggs)
	Router.Route("dj", `Send the song "Goudja On The Floor - Risitas ft. Cocinero Dumas".`, Router.Eggs)
}

func main() {
	log.Println("==================================================")
	log.Printf("%s - Version %s - PID %d\n", os.Getenv("BOT_TITLE"), os.Getenv("BOT_VERSION"), os.Getpid())

	// Creates a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("[ERROR] Could not create Discord session: %v", err)
	}

	// Registers the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(Router.OnMessageCreate)

	// Only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Opens a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Fatalf("[ERROR] Could not open connection: %v", err)
	}

	// Loads all data
	err = data.Load(os.Getenv("DATA_PATH"))
	if err != nil {
		log.Fatalf("[ERROR] Could not load data: %v", err)
	}

	// Updates bot status every X minutes with random activity
	min, _ := strconv.Atoi(os.Getenv("UPDATE_ACTIVITY_DURATION"))
	go AutoUpdateStatus(dg, min)

	// Notifies each guild of the release of a new video
	min, _ = strconv.Atoi(os.Getenv("CHECK_NEW_VIDEO_DURATION"))
	go NotifyNewVideo(dg, min)

	// Waits here until CTRL-C or other term signal is received
	log.Println("[INFO] Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

// autoUpdateStatus picks a random activity from data file
// and applies it to the Discord bot status every X minutes.
func AutoUpdateStatus(dg *discordgo.Session, min int) {
	activities := data.Activities()
	if len(activities) == 0 {
		log.Fatalf("[ERROR] Array must contain at least one activity. got=%v", len(activities))
	}

	// Picks the first activity for first running
	dg.UpdateGameStatus(0, activities[0])

	// Picks a random activity each X minutes
	duration, _ := time.ParseDuration(fmt.Sprintf("%dm", min))
	for range time.Tick(duration) {
		log.Println("[INFO] AutoUpdateStatus called")

		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(activities))

		dg.UpdateGameStatus(0, activities[i])
	}
}

// notifyNewVideo checks every X minutes if a new video of Chef Michel Dumas
// is released and notifies all guilds connected to the bot.
func NotifyNewVideo(dg *discordgo.Session, min int) {
	var lastVideoLink string
	fp := gofeed.NewParser()

	// Checks if there is a new video each X minutes
	duration, _ := time.ParseDuration(fmt.Sprintf("%dm", min))
	for range time.Tick(duration) {
		log.Println("[INFO] NotifyNewVideo called")

		feed, _ := fp.ParseURL("https://www.youtube.com/feeds/videos.xml?channel_id=UCSLyEx8ISkp567AjOAHYN5Q")

		video := feed.Items[0]
		shortLink := "https://youtu.be/" + strings.Split(video.Link, "=")[1]

		if lastVideoLink == "" {
			// Prevents notification when the bot starts and updates cache
			lastVideoLink = shortLink
		} else if lastVideoLink != shortLink {
			// Updates the cache
			lastVideoLink = shortLink

			// Sends the video notification to all servers having the bot
			msg := "Nouvelle vidéo, comme ça !\n" + shortLink
			SendMessageAllGuilds(dg, msg)
		}
	}
}

// SendMessageAllGuilds sends a message to all guilds connected to the bot.
// The bot will send a message in the first textual channel it finds of each guild.
func SendMessageAllGuilds(dg *discordgo.Session, msg string) {
	log.Println("[INFO] SendMessageAllGuilds called")

	// Loop through each guild in the session
	for _, g := range dg.State.Guilds {
		// Get channels for this guild
		channels, _ := dg.GuildChannels(g.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Now, sends a message only to the first channel found
			// Later on, we can select the channel for notifications
			dg.ChannelMessageSend(c.ID, msg)
			// fmt.Sprintf("testmsg (sorry for spam). Channel name is %q", c.Name),
			return
		}
	}
}
