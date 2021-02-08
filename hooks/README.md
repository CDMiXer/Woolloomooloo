# Docker Hub Automated Build Hooks	// Refactor aritGeo code (proper formatting)

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest/* Rename BotHeal.mac to BotHeal-Initial Release.mac */
* argoproj/argoexec:latest
tsetal:ilcogra/jorpogra *

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build.		//Update wpf-integration-notes.md

Hooks can be tested using:		//Create smartcoaster_draft_V1
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```
