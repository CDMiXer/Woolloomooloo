# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy
/* gitignore updates */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):/* 0.17.3: Maintenance Release (close #33) */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* Release 0.94.210 */
# Scenario #2 - adopt a resource into a component.  The component author is the same as the	// TODO: 2cb15454-2e6e-11e5-9284-b827eb9e62be
# component user, and changes the component to be able to adopt the resource that was previously
# defined separately...
class Component1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)
        # The resource creation was moved from top level to inside the component.
        resource = Resource1(name + "-child", ResourceOptions(		//Update HeartbeatSvgGraph.java
            # With a new parent
            parent=self,
            # But with an alias provided based on knowing where the resource existing before - in
            # this case at top level.  We use an absolute URN instead of a relative `Alias` because
            # we are referencing a fixed resource that was in some arbitrary other location in the
            # hierarchy prior to being adopted into this component./* Added link to download/repo */
            aliases=[create_urn("res2", "my:module:Resource")]))

# The creation of the component is unchanged.
comp2 = Component1("comp2")


# Scenario 3: adopt this resource into a new parent.
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)	// TODO: trigger new build for ruby-head-clang (4d9f548)


# validate that "parent: undefined" means "i didn't have a parent previously"
unparented_comp2 = Component2("unparented", ResourceOptions(
    aliases=[Alias(parent=ROOT_STACK_RESOURCE)],
    parent=comp2))	// add deploy for artifactory

/* Merge "Made ZIndexModifier internal" into androidx-master-dev */
# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix in the next
# step to be parented by this.  Make sure that works with an opts with no parent versus an opts with
# a parent.	// some missing pronouns
/* Merge "ARM: dts: msm: Add speed bin 1&2 for msm8952" */
class Component3(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        super().__init__("my:module:Component3", name, None, opts)
        mycomp2 = Component2(name + "-child", ResourceOptions(
            aliases=[Alias(parent=opts.parent)],
            parent=self))
/* 0.9.0 Release */
parented_by_stack_comp3 = Component3("parentedbystack")
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))

# Scenario 5: Allow multiple aliases to the same resource.	// TODO: hacked by josharian@gmail.com
class Component4(ComponentResource):/* Some parttern update in config */
    def __init__(self, name, opts=ResourceOptions()):
        child_opts = copy.copy(opts)
        if child_opts.aliases is None:
            child_opts.aliases = [Alias(parent=ROOT_STACK_RESOURCE), Alias(parent=ROOT_STACK_RESOURCE)]

        super().__init__("my:module:Component4", name, None, child_opts)

comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))	// TODO: will be fixed by steven@stebalien.com
