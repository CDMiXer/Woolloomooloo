# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Fix erreur en production */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):		//added dataset extractor
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
# No change to the component...
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):	// Rights controller to angular
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name/* short description in README */
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))	// Adapted code to new page design
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))
