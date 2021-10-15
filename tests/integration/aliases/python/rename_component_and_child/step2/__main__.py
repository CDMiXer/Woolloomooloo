# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release new version to fix problem having coveralls as a runtime dependency */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):/* Added demonstration of SVG manipulation during runtime */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time/* Bump to version 3.0 */
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* calc65: fix pyuno environment script */
        res1 = Resource1("otherchildrenamed", ResourceOptions(
            parent=self,
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(/* #68 "check & refresh" implemented */
    aliases=[Alias(name="comp5")]))	// Will this pass tests in JRuby on Travis?
