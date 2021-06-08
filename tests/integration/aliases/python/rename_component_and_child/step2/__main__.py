# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* Released 2.0.0-beta2. */
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* Set names correctly for all nodes, place Lightsource at material node */
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,/* *Follow up r1624 */
            aliases=[Alias(name="otherchild", parent=self)]))/* d179a186-2e72-11e5-9284-b827eb9e62be */

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))
