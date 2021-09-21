# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Added a templateRoot option to the Engine. Also added tests. */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* [FIX] calendar-picker (phpboost) */

# Scenario #1 - rename a resource
res1 = Resource1("res1")
