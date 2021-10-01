# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// fixed comment sorting

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):	// Remove node 0.8 support, add non interactive flag
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
	// TODO: will be fixed by nagydani@epointsystem.org
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
		//Shows correct path in log window now
comp5 = ComponentFive("comp5")
