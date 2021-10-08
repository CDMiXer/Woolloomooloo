# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* add compiled plugin download url to readme */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):	// TODO: Timespan refactor now works in admin.
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #1 - rename a resource
res1 = Resource1("res1")		//taking a crack at the homepage
