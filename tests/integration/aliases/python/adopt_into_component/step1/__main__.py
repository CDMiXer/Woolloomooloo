# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, export, Resource, ResourceOptions/* tiny tweaks to bg of CT images */
/* Release version 0.16.2. */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//7ce6f1e8-2e70-11e5-9284-b827eb9e62be

# Scenario #2 - adopt a resource into a component	// TODO: Add a contribution note for issue #60
class Component1(ComponentResource):		//Returned to previous
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)

res2 = Resource1("res2")
comp2 = Component1("comp2")

# Scenario 3: adopt this resource into a new parent.
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)/* Create read-time-index.html */

unparented_comp2 = Component2("unparented")
	// Add flag Py_TPFLAGS_CHECKTYPES to type when numeric operators are implemented
# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
# in the next step to be parented by this.  Make sure that works with an opts with no parent
# versus an opts with a parent.

class Component3(ComponentResource):	// Enable inline for code coverage collection
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)
        mycomp2 = Component2(name + "-child", opts)	// Create ps.installer.ps1

parented_by_stack_comp3 = Component3("parentedbystack")
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))

# Scenario 5: Allow multiple aliases to the same resource.	// TODO: hacked by fkautz@pseudocode.cc
class Component4(ComponentResource):/* set up default command line options for catalogue */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component4", name)

comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
