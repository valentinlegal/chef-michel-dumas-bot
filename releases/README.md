# Releases

This folder will contain all releases generated locally via the `make release` command.

The [update.sh](./update.sh) script will download the desired release from the GitHub repository and install it. It will not be copied to the release archive.

The [install.sh](./install.sh) script will be copied into the release archive and will allow the installation of the bot.

Releases are not versioned.

## Generate a release

```sh
make release
```

## Install or update a release in prod environment

- For a first installation, download the [update.sh](./releases/update.sh) file to your production server:

```sh
curl -LO https://raw.githubusercontent.com/valentinlegal/chef-michel-dumas-bot/main/releases/update.sh
chmod +x update.sh
```

- Run the following command to install or update the bot:

```sh
./update.sh <VERSION>

# e.g. ./update.sh 1.0.0
```

- Go to the `app` directory:

```sh
cd app/
```

- Check and complete the `.env` file.

- Run the bot:

```sh
./chef-michel-dumas-bot

# In background with logs file
./chef-michel-dumas-bot >> app.log 2>&1 &
```
