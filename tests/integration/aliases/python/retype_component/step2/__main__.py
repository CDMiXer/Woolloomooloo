# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* Released GoogleApis v0.1.0 */
        super().__init__("my:module:Resource", name, None, opts)
/* 6ab75cbe-2e71-11e5-9284-b827eb9e62be */
# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:/* Release version: 1.9.1 */
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)
	// TODO: will be fixed by igor@soramitsu.co.jp
        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
