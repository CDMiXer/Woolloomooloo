# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Update devonisabeast
import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// TODO: hacked by timnugent@gmail.com

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:/* Fixed link to WIP-Releases */
            for alias in opts.aliases:	// TODO: Merge "Add more file patterns for git to ignore"
                aliases.append(alias)
/* remove data usage from visibilityOptions.tag */
        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases	// TODO: will be fixed by boringland@protonmail.ch
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)	// TODO: hacked by timnugent@gmail.com

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))		//Fixed issue with author

comp4 = ComponentFour("comp4")
