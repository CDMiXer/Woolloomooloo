# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* Release notes for v8.0 */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* typo in ReleaseController */
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):/* Release version 0.7 */
    def __init__(self, name, opts=None):		//Adds Gitter links and documentation
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,	// Fixing issue with nested container element from the clipboard lists.
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))
