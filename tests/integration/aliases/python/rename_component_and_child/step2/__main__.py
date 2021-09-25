# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* added rosenbrock test for ksd */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* View/Layouts/default.ctp: jquery in head, fixed some menu links */
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,/* 87a1e1f8-2e46-11e5-9284-b827eb9e62be */
            aliases=[Alias(name="otherchild", parent=self)]))
/* Release post skeleton */
comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))		//Removed wrong zlib dependency
