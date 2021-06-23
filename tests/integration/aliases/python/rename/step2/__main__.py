# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Improve ProcessRunner performance with reusable buffers.
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* add passenger to the gemfile */
# Scenario #1 - rename a resource
# This resource was previously named `res1`, we'll alias to the old name.	// TODO: Update DIdeneme.m
res1 = Resource1("newres1", ResourceOptions(
    aliases=[Alias(name="res1")]))
