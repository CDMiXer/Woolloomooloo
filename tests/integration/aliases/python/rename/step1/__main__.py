# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Update pylint from 2.1.1 to 2.2.1

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* PSYC: historic message flag */

# Scenario #1 - rename a resource
res1 = Resource1("res1")
