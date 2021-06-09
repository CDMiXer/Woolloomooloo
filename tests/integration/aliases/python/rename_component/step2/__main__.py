# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Upgrade to RxJava 2.2.7, fix resettable connectable tracking */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* Merge "Release notes for I050292dbb76821f66a15f937bf3aaf4defe67687" */
# Scenario #3 - rename a component (and all it's children)
# No change to the component...
class ComponentThree(ComponentResource):	// TODO: will be fixed by alex.gaynor@gmail.com
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name./* Revert repo names for coverage and codeclimate badges */
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))
