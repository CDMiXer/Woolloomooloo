# Docker Hub Automated Build Hooks	// TODO: hacked by lexy8russo@outlook.com

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest/* Added 2 safe checks to avoid PHP notices. */
/* Release 0.32.1 */
It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build./* Version 2.1.0 Release */

Hooks can be tested using:
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```
