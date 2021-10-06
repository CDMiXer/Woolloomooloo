# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Implement support for "Matplotlib" 3.3.0.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: GameUI test started
/* (vila) Release 2.3.3 (Vincent Ladeuil) */
class Resource1(ComponentResource):/* Release of eeacms/www:20.9.13 */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
# No change to the component...
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))
		//Prevent direct access for .blade.php files
# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(/* first release; M1 */
    aliases=[Alias(name="comp3")]))
