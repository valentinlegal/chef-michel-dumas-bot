#!/bin/bash

# ##################################################
# Script Name : update.sh
# Description : Downloads the desired release from the GitHub repository and installs it.
# Args        : The version to download.
# Usage       : ./update.sh <VERSION>
# Notes       : The install.sh script is required.
# Version     : 1.0.0
# Author      : Valentin Le Gal
# ##################################################

REPOSITORY_AUTHOR=valentinlegal
REPOSITORY_NAME=chef-michel-dumas-bot
VERSION=$1
FILE_NAME=$REPOSITORY_NAME-$VERSION.tar.gz
FILE_URL=https://github.com/$REPOSITORY_AUTHOR/$REPOSITORY_NAME/releases/download/v$VERSION/$FILE_NAME

# Checks if version is supplied
if [ -z $VERSION ]
then
    echo -e "\033[0;31m[ERROR]\033[m No version supplied (e.g. 1.0.0)"
    exit 1
fi

# Download the release if not exists
if ! [ -f $FILE_NAME ]
then
  curl -LO $FILE_URL -f
  if [ $? -ne 0 ]
  then
    echo -e "\033[0;31m[ERROR]\033[m Version $VERSION doesn't exist"
    exit 1
  fi
else
  echo "$FILE_NAME already exists: download cancelled"
fi

# Extract the release in a clean "app" folder
rm -rf app/
tar -xf $FILE_NAME

# Runs the installation / update script
cd app/ && ./install.sh

echo -e "\033[0;32m[OK]\033[m The bot is up to date!"
echo -e "Next steps:"
echo -e "  1. cd app/"
echo -e "  2. vi .env"
echo -e "  3. ./chef-michel-dumas-bot"
