# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Release for 24.6.0 */
/* Release 0.0.4 maintenance branch */
# Scenario #1 - rename a resource
# This resource was previously named `res1`, we'll alias to the old name./* Finalization of v2.0. Release */
res1 = Resource1("newres1", ResourceOptions(	// TODO: will be fixed by mikeal.rogers@gmail.com
    aliases=[Alias(name="res1")]))
