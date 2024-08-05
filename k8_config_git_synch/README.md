# File update through git repository on k8s

Repository for the file https://github.com/sdrubolo/songs/blob/master/albums.json

This folder contains:

- a git-sync which will look for updates of the git repository
- a simple app which will return the content of the file

## To run script container

`helm install go-config . --namespace go-config --create-namespace --wait`

## To run app container

`helm install go-config-app . --namespace go-config`