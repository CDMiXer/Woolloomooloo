# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 2.0.0-rc.17 */
/* Added standard cost to item's details page. */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Cleanup and ReleaseClipX slight fix */

# Scenario #3 - rename a component (and all it's children)	// Updated readme based on further state in project
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

comp3 = ComponentThree("comp3")
