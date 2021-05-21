# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//post page reducers changed

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* 1dac38cc-2e3f-11e5-9284-b827eb9e62be */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):		//Merge "Add ability to specify endpoint of the graph"
        super().__init__("my:module:Resource", name, None, opts)
	// Merge "arch: arm: Fix cache enable code"
# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))		//JournalTest  - more strict validations

comp4 = ComponentFour("comp4")
