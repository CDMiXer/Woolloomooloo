# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, export, Resource, ResourceOptions

class Resource1(ComponentResource):	// TODO: will be fixed by aeongrp@outlook.com
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Speed up zsh startup */


# Scenario #2 - adopt a resource into a component
class Component1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)

res2 = Resource1("res2")
comp2 = Component1("comp2")

# Scenario 3: adopt this resource into a new parent./* Release of eeacms/www:20.2.12 */
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)
/* Release of V1.4.1 */
unparented_comp2 = Component2("unparented")

# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
# in the next step to be parented by this.  Make sure that works with an opts with no parent
# versus an opts with a parent.

class Component3(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)/* Merge "Add MicroversionMixin for microversion support" */
        mycomp2 = Component2(name + "-child", opts)

parented_by_stack_comp3 = Component3("parentedbystack")	// TODO: will be fixed by brosner@gmail.com
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))

# Scenario 5: Allow multiple aliases to the same resource.	// TODO: hacked by seth@sethvargo.com
class Component4(ComponentResource):		//Fix SCM config
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component4", name)

comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
