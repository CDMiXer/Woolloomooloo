# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Added pruebaTecnica.xml
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Delete Release 3.7-4.png */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time	// TODO: hacked by greg@colvin.org
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchildrenamed", ResourceOptions(		//DiskFilesystem needs unistd.h
            parent=self,
            aliases=[Alias(name="otherchild", parent=self)]))

comp5 = ComponentFive("newcomp5", ResourceOptions(
    aliases=[Alias(name="comp5")]))
