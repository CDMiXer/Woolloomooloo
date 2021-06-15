# Docker Hub Automated Build Hooks
		//Fix i18n FR typo
This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build./* Release version 2.0.0.RC2 */

Hooks can be tested using:
```/* Release 33.4.2 */
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```
