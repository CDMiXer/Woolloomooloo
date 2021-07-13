# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Merge branch 'master' into greenkeeper/tape-4.6.3
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):	// https://pt.stackoverflow.com/q/56830/101
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #1 - rename a resource
res1 = Resource1("res1")
