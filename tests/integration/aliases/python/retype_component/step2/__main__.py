# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: hacked by steven@stebalien.com

class Resource1(ComponentResource):/* Merge "msm: clock-8960: Unstick usb_hsic_hsic_clk halt bits" into msm-3.0 */
    def __init__(self, name, opts=None):/* fix broken translation for the 'welcome back' message in the lobby view */
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):		//Merge "Adding drag outlines for dragging out of folders"
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)		//Some python exports for handling music stuff.

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.	// TODO: hacked by sjors@sprovoost.nl
        res1 = Resource1("otherchild", ResourceOptions(parent=self))		//Restore eof line.

comp4 = ComponentFour("comp4")
