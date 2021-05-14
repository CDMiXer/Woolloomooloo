# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//merge logging_gearman build fix
# Scenario #1 - rename a resource
# This resource was previously named `res1`, we'll alias to the old name.
res1 = Resource1("newres1", ResourceOptions(/* Create 186.Reverse Words in a String II.md */
    aliases=[Alias(name="res1")]))/* Released version 0.6.0 */
