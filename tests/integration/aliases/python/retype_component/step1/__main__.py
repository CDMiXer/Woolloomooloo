# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Merge devYears into development
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// Create asteroid.py

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):/* Merge "Release Notes 6.0 -- Other issues" */
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
