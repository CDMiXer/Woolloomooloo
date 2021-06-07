# Docker Hub Automated Build Hooks

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.		//forgot to update the gemspec date
This is needed since we publish multiple images as part of a single build:/* Update libraries/legacy/module/helper.php */
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during/* try to keep data in browser */
the build.		//keine Gewalt

Hooks can be tested using:
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```
