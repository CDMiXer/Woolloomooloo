# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* Create newsandupdate.css */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* Release 1.6.1 */
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):	// 830d25d2-2e53-11e5-9284-b827eb9e62be
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
