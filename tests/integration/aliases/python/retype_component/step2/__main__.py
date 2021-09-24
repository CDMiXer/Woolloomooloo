# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Delete sitesearch.xml

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// TODO: hacked by witek@enjin.io

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)	// TODO: hacked by juan@benet.ai

        # ..and then make the super call with the new type of this resource and the added alias./* Released 2.2.4 */
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.	// TODO: will be fixed by arajasek94@gmail.com
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")	// TODO: Fixed chain rule of LMA.
