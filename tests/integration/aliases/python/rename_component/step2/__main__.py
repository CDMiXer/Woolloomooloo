# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge "Add new API to Animator to allow seeking of animations" */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)		//Added formatter for component diagram
# No change to the component...	// TODO: PHP 5.4 compliant
class ComponentThree(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the/* Update packages/page-alerts/CHANGELOG.md */
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name.
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))
/* Put BLAS calls in VPolyFit, but doesn't give correct answer yet. */
# ...but applying an alias to the instance successfully renames both the component and the children.
comp3 = ComponentThree("newcomp3", ResourceOptions(
    aliases=[Alias(name="comp3")]))
