# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Release 4.0.10.54 QCACLD WLAN Driver" */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* a7674658-2e56-11e5-9284-b827eb9e62be */

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))	// TODO: Added jruby script add-on

comp5 = ComponentFive("comp5")
