#!/bin/sh

# Clone the repository if not already present
if [ ! -d "/songs/.git" ]; then
  git clone $REPO_NAME /songs
fi

cd "/songs"

# Periodically pull the latest changes
while true; do

  git pull origin master
  sleep 30 # Wait for 30 seconds before pulling again
done