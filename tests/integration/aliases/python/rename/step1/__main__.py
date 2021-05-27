# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release version 3.2.0 build 5140 */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* 9e37b9a0-2e74-11e5-9284-b827eb9e62be */
class Resource1(ComponentResource):/* Merge "all: Windows fixes (don't listen on file descriptors in test.World, etc)" */
    def __init__(self, name, opts=None):/* Release "1.1-SNAPSHOT" */
        super().__init__("my:module:Resource", name, None, opts)
/* rebar magick in app */
# Scenario #1 - rename a resource
res1 = Resource1("res1")
