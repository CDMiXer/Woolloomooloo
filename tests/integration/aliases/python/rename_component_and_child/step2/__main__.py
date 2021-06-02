# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Add Evorus news cover

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):/* [Doc] Change classname from DoctrineConverter to DoctrineParamConverter */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):/* dae3f4e4-2e66-11e5-9284-b827eb9e62be */
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchildrenamed", ResourceOptions(		//Timestamp for the private user file
            parent=self,
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))
