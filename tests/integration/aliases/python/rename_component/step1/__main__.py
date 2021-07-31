# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):		//Removed pow for integer exponents.
        super().__init__("my:module:ComponentThree", name, None, opts)/* Issue #512 Implemented MkReleaseAsset */
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name/* Preparing Changelog for Release */
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))
	// TODO: will be fixed by xiemengjun@gmail.com
comp3 = ComponentThree("comp3")		//device density
