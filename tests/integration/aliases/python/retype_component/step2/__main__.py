# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release of eeacms/jenkins-master:2.235.2 */

import copy
	// this file was missing preventing manual build
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
/* Release v0.5.3 */
# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):/* spaced readme */
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]/* Merged QA into master */
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)
/* Added dragging tests to solve #147 */
        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.		//Create Floyd-Warshall Algorithm
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
