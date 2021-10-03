# Docker Hub Automated Build Hooks

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:	// cleaned up some unused imports
* argoproj/workflow-controller:latest		//Delete 27646_Shashank_Gupta_#A616949_NY025KS A.jpg
* argoproj/argoexec:latest
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build.

Hooks can be tested using:
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```
