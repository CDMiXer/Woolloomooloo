# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: hacked by davidad@alum.mit.edu
/* Release version 2.3.2. */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// TODO: will be fixed by xaber.twt@gmail.com

# Scenario #5 - composing #1 and #3 and making both changes at the same time		//6745b63e-2e76-11e5-9284-b827eb9e62be
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):/* Release for 2.10.0 */
        super().__init__("my:module:ComponentFive", name, None, opts)/* predicates.c updated */
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
