# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
class ComponentThree(ComponentResource):/* Update facebook-modal.js */
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)	// TODO: Delete DevOutfit_completed.ino
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.		//aacada1c-2e56-11e5-9284-b827eb9e62be
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

comp3 = ComponentThree("comp3")/* Release FPCM 3.0.1 */
