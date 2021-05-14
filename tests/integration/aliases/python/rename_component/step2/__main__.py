# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
		//Update exit.js
class Resource1(ComponentResource):	// fixes #1681 on source:branches/1.11
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
# No change to the component...
class ComponentThree(ComponentResource):/* Created 0xrnair_monster_chicken_burrito.md */
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name./* [Translating] Guake 0.7.0 Released – A Drop-Down Terminal for Gnome Desktops */
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))	// rev 537785

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(/* Release of eeacms/www:20.11.26 */
    aliases=[Alias(name="comp3")]))
