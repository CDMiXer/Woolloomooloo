# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Publishing post - Why I Started Learning Programming */

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component/* [artifactory-release] Release version 1.4.0.RELEASE */
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases	// Update WTracker/WTracker.m
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)

        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.		//improved signal control
        res1 = Resource1("otherchild", ResourceOptions(parent=self))/* [Release] sticky-root-1.8-SNAPSHOTprepare for next development iteration */

comp4 = ComponentFour("comp4")
