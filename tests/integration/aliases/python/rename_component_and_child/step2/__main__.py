# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Delete imagenaming.cpp.save.1
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* Release of eeacms/eprtr-frontend:0.2-beta.14 */
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))	// TODO: will be fixed by hugomrdias@gmail.com
