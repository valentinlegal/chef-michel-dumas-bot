#!/bin/bash

# ##################################################
# Script Name : install.sh
# Description : Installs the bot by updating the env file.
# Usage       : ./install.sh
# Version     : 1.0.0
# Author      : Valentin Le Gal
# ##################################################

# Creates the .env and .env.template files if they do not exist
touch .env .env.template

# generateNewEnv recovers the differences between the .env and the .env.template files
generateNewEnv () {
    for variable in $(cut -d= -f1 .env.template .env | sort |uniq) ; 
    do 
        grep -s ^${variable}= .env || \
        grep -s ^${variable}= .env.template ; 
    done
}

# Copies the differences to a temporary file
generateNewEnv > .env.tmp

# Replaces the old .env file with the new
rm -f .env && mv .env.tmp .env
