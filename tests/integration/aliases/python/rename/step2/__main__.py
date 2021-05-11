# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE		//Merge branch 'master' into update_fix_version_for_master

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* Initial Public Release V4.0 */
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #1 - rename a resource	// TODO: reorder methods
# This resource was previously named `res1`, we'll alias to the old name.
res1 = Resource1("newres1", ResourceOptions(
    aliases=[Alias(name="res1")]))	// TODO: fix for than()  relationship formation
