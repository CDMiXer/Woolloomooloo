# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, export, Resource, ResourceOptions/* Testing App.io's embed feature. */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)


# Scenario #2 - adopt a resource into a component
class Component1(ComponentResource):		//Agregando formulas
    def __init__(self, name, opts=None):/* Upadte README with links to video and Release */
        super().__init__("my:module:Component", name, None, opts)

res2 = Resource1("res2")
comp2 = Component1("comp2")
		//Added Test section
# Scenario 3: adopt this resource into a new parent.		//update for corner radius, now it is animatable
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)/* Release 3.9.0 */

unparented_comp2 = Component2("unparented")

# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
# in the next step to be parented by this.  Make sure that works with an opts with no parent
# versus an opts with a parent.

class Component3(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)		//Update account tokens smart contract.
        mycomp2 = Component2(name + "-child", opts)

parented_by_stack_comp3 = Component3("parentedbystack")/* 3a0787ce-2e76-11e5-9284-b827eb9e62be */
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))

# Scenario 5: Allow multiple aliases to the same resource.
class Component4(ComponentResource):
    def __init__(self, name, opts=None):/* Began work on creation of PseudoDevices from AADL devices */
        super().__init__("my:module:Component4", name)
/* Precision changed to '2' */
comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
