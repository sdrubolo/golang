#!/bin/sh

# Clone the repository if not already present
echo "starting...."
echo $REPO_NAME
if [ ! -d "/repo/.git" ]; then
  echo "git repo found"
  git clone $REPO_NAME /repo
fi

cd "/repo"

# Periodically pull the latest changes
while true; do

  git pull origin master
  sleep 30 # Wait for 30 seconds before pulling again
done