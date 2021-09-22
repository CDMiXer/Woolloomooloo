# Docker Hub Automated Build Hooks

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:		//New post: Perma01test
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest	// Automatic changelog generation for PR #49236 [ci skip]
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during/* Fixed checkstyle errors for overloaded method order */
the build.

Hooks can be tested using:
```	// Add support for event contexts
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```	// TODO: Update WinnersGroovyAndJavaMixed.groovy
