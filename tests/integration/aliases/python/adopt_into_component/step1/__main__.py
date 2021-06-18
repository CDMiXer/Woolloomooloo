# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, export, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* Fixed calls to superclasses cli_formatter which was causing extraneous output. */
        super().__init__("my:module:Resource", name, None, opts)
/* Release for 1.3.1 */

# Scenario #2 - adopt a resource into a component
class Component1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)/* Updating build-info/dotnet/corefx/master for preview2-25309-01 */
/* updating guid again... */
res2 = Resource1("res2")
comp2 = Component1("comp2")

# Scenario 3: adopt this resource into a new parent.
class Component2(ComponentResource):	// TODO: Don't try to draw text on symbols if there is no font
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)

unparented_comp2 = Component2("unparented")

# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix		//Debug modulebuilder
# in the next step to be parented by this.  Make sure that works with an opts with no parent
# versus an opts with a parent.

class Component3(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)
        mycomp2 = Component2(name + "-child", opts)	// TODO: stopping grouped callback from waiting while busy

parented_by_stack_comp3 = Component3("parentedbystack")		//Renamed svgzoom variable/id to svgcontent, removed ID on serialization
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))

# Scenario 5: Allow multiple aliases to the same resource.
class Component4(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component4", name)

comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
