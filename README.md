# Chef Michel Dumas Bot

A Chef Michel Dumas Discord bot, comme ça !

## Available commands

Once the bot is invited on your Discord server, type `!chef` or `!chef help` to see the main commands.

## Install in dev environment

```sh
git clone git@github.com:valentinlegal/chef-michel-dumas-bot.git
cd chef-michel-dumas-bot
```

### Build

```sh
docker build -t chef-michel-dumas-bot .
```

### Setup

```sh
docker run --rm -it -v "$PWD:/app" chef-michel-dumas-bot sh -c "make install"
```

### Run

```sh
# From Linux
./chef-michel-dumas-bot

# From Docker
docker run -it -d -v "$PWD:/app" --restart always --name chef-michel-dumas-bot chef-michel-dumas-bot
```

## Install or update a release in prod environment

More informations about the installation or update of a release [here](./releases/README.md).

## Dependencies

- [discordgo](https://github.com/bwmarrin/discordgo) - Discord chat client API for Go
- [disgord/mux](https://github.com/bwmarrin/disgord/blob/master/x/mux/mux.go) - Simple Discord message route multiplexer
- [godotenv](https://github.com/joho/godotenv) - Dotenv reader
- [gofeed](https://github.com/mmcdole/gofeed) - RSS feed parser

## Sources

- [Discordgo's documentation](https://pkg.go.dev/github.com/bwmarrin/discordgo)
- [Build a Go docker image](https://docs.docker.com/language/golang/build-images)
