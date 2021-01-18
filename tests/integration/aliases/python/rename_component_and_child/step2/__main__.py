# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//AJAX owner

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
		//using the stored user name
class Resource1(ComponentResource):		//Update the event demo.
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//Update data_gathering_methods.md
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,/* Starting mediaqueries for grid system */
            aliases=[Alias(name="otherchild", parent=self)]))	// Refactored Tree and HashTree

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))
