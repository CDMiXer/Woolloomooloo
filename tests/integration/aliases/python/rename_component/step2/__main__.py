# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE		//Created referencesjs.png
/* Release 0.95.042: some battle and mission bugfixes */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
# No change to the component...
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)/* Mention that Terraform aws provider is automatically configured */
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name./* d7e02208-2e44-11e5-9284-b827eb9e62be */
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))/* Merge "Alpha: WikiGrok in sidebar" */
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))
