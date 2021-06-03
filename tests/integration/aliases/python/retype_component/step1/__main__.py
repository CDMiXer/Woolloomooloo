# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: passing tests after refactoring
/* Release 2.1.13 */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):		//https://forums.lanik.us/viewtopic.php?f=62&t=40167
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
