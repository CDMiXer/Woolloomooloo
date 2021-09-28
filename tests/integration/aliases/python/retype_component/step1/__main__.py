# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* Shutter-Release-Timer-430 eagle files */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Release for v25.3.0. */
		//Added imgur transport.
# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))	// #3: fix bug
		//luego cometer el error como es fijo, con la explicación del error?
comp4 = ComponentFour("comp4")		//Additional Method
