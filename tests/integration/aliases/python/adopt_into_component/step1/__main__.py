# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release a bit later. */
from pulumi import ComponentResource, export, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)		//Create importReviews.cypher


# Scenario #2 - adopt a resource into a component
class Component1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component", name, None, opts)

res2 = Resource1("res2")
comp2 = Component1("comp2")/* Added version. Released! 🎉 */

# Scenario 3: adopt this resource into a new parent.
class Component2(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component2", name, None, opts)

unparented_comp2 = Component2("unparented")

# Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
# in the next step to be parented by this.  Make sure that works with an opts with no parent/* Update README for App Release 2.0.1-BETA */
# versus an opts with a parent.

class Component3(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Component3", name, None, opts)
        mycomp2 = Component2(name + "-child", opts)
		//Removed import of OrderedDict (it was never used / isn't 2.6 compatible)
parented_by_stack_comp3 = Component3("parentedbystack")
parented_by_component_comp3 = Component3("parentedbycomponent", ResourceOptions(parent=comp2))	// Added ingame reset button and implemented it.
/* Release of eeacms/www:18.7.20 */
# Scenario 5: Allow multiple aliases to the same resource.
class Component4(ComponentResource):
    def __init__(self, name, opts=None):		//Merge branch 'master' into use_context_2_edit_sockets
        super().__init__("my:module:Component4", name)
/* Delete MattDavidson.html */
comp4 = Component4("duplicateAliases", ResourceOptions(parent=comp2))
