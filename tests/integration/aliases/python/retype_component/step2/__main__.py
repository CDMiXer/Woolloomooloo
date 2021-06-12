# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)		//Fix README paragraph boundary

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):/* Added alternative subject form of the pronoun hen, analysis only. */
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.	// TODO: stub rpc servers
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented/* Added null checks to oldState->Release in OutputMergerWrapper. Fixes issue 536. */
        # to.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")		//Merge "Upgrade the rst convention of the Admin Guide"
