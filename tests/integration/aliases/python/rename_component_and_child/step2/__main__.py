# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):/* a646e314-4b19-11e5-897b-6c40088e03e4 */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// TODO: Renommage Key

# Scenario #5 - composing #1 and #3 and making both changes at the same time/* Allow singpath problem to be reset  */
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* Updated docs to reflect changes to project layout. */
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))	// Switch to variable width nodes
