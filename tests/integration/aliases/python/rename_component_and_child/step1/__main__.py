# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 2.7.3 */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
:)enoN=stpo ,eman ,fles(__tini__ fed    
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))	// TODO: Update gnmapParse.py
/* Backplotting wasn't working for me in Windows, since installing python 2.6 */
comp5 = ComponentFive("comp5")
