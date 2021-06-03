# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// fixing theme

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Update RFC0013-PowerShellGet-PowerShellGallery_PreRelease_Version_Support.md */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
# No change to the component...	// TODO: will be fixed by xaber.twt@gmail.com
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)		//bundle-size: 558439d97cd0ab09c0b979e1a55516346a2c2b3c.json
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))

# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))
