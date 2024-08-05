# File update through git repository

Repository for the file https://github.com/sdrubolo/songs/blob/master/albums.json

This folder contains:

- a script which will look for updates of the git repository
- a simple app which will return the content of the file

## To run script container

`VERSION=0.1 REPO_NAME=<GIT_URL_FOR_CLONE> REPO_PATH=<INTERNAL_PATH> docker-compose up --build`

## To run app container

`VERSION=0.1 SONGS_PATH=<PATH_TO_GITHUB_REPO_DIR> FILE_PATH=<INTERNAL_PATH> docker-compose up --build`