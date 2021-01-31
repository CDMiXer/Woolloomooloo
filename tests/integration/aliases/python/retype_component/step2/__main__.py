# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 11b3c26a-2e49-11e5-9284-b827eb9e62be */

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):/* Changed the Changelog message. Hope it works. #Release */
    def __init__(self, name, opts=ResourceOptions()):		//enabled parameter flipping for the xss module
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:/* Release Alpha 0.6 */
            for alias in opts.aliases:/* Release 1.0.27 */
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases/* Use Releases to resolve latest major version for packages */
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
/* Ensure item feedback text is shown for all view types. */
comp4 = ComponentFour("comp4")
