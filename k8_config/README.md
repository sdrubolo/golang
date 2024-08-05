# File update through git repository on k8s

Repository for the file https://github.com/sdrubolo/songs/blob/master/albums.json

This folder contains:

- a script which will look for updates of the git repository
- a simple app which will return the content of the file

# Build App

Naviage to /app folder and run `VERSION=0.1 docker-compose build`

# Build Script

Naviage to /git_watcher folder and run `VERSION=0.1 docker-compose build`

## To run script container

`helm install go-config . --namespace go-config --create-namespace --wait`

## To run app container

`helm install go-config-app . --namespace go-config`