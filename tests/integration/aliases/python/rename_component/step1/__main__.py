# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* [artifactory-release] Release version 3.2.17.RELEASE */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE/* Release 2.2.10 */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #3 - rename a component (and all it's children)
class ComponentThree(ComponentResource):/* changed 0 to 1 :P */
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentThree", name, None, opts)
        # Note that both un-prefixed and parent-name-prefixed child names are supported. For the
        # later, the implicit alias inherited from the parent alias will include replacing the name
        # prefix to match the parent alias name./* Update readme to Adldap2 v7 */
        resource1 = Resource1(name + "-child", ResourceOptions(parent=self))
        resource2 = Resource1("otherchild", ResourceOptions(parent=self))
		//New version of Xodogo - 1.1.1
comp3 = ComponentThree("comp3")
