# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):		//14d746ba-2e4e-11e5-9284-b827eb9e62be
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #1 - rename a resource
res1 = Resource1("res1")
