# Docker Hub Automated Build Hooks	// TODO: hacked by souzau@yandex.com
/* HT.Hexagon.Id attribute is now lowercase */
This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest/* Added FsprgEmbeddedStore/Release, Release and Debug to gitignore. */

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build.
		//Release 0.13.4 (#746)
Hooks can be tested using:
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push	// TODO: will be fixed by sebastian.tharakan97@gmail.com
```/* [MIN] GUI, InfoView: only refresh text if view is visible */
