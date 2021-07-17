# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
		//Fix Bug 3861 - Phenotype filter - Inference not working
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* Added new functions to start/stop productionsites to LuaMap and fixed the test. */
        super().__init__("my:module:Resource", name, None, opts)/* Merge "Release notes: prelude items should not have a - (aka bullet)" */

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
	// benerin search saran humas
comp4 = ComponentFour("comp4")
