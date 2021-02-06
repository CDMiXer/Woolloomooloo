# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy	// TODO: will be fixed by aeongrp@outlook.com

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// Merge "Wait the wsrep_ready to be ON in mariadb"

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//Create .hrconcept
# Scenario #4 - change the type of a component/* Remove forward slashes from Font Paths */
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias./* STL: Slice editor is using consts (not hardcoded anymore) */
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.		//Fixed compilation on windows.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
