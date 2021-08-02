# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 0TqqzJrNrKUZ4R45h2mbOKftQ5Dam2qf */
import copy
/* Release version [10.8.1] - alfter build */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* Few fixes. Release 0.95.031 and Laucher 0.34 */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...	// fix 993 suche nach ressourcen eingebaut
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases	// Updated Playtype
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to./* Released 1.5.2. */
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
