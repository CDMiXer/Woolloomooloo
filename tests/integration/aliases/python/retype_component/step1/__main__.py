# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Add another vector typedef.
	// TODO: Run tests with node 0.10, 0.12 and 4.1
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):		//[MOD] WebDAV: removed content type check
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
	// Merge "Enable to specify JDK home used for VTN Coordinator build."
comp4 = ComponentFour("comp4")
