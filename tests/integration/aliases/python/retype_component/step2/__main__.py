# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import copy

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: Restricted users table to superuser only

class Resource1(ComponentResource):		//Generate name from id
    def __init__(self, name, opts=None):/* Merge "Add spec file for receiver action module" */
        super().__init__("my:module:Resource", name, None, opts)
		//Create problem14.py
# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=ResourceOptions()):/* Use FaradayMiddleware::Mashify */
        # Add an alias that references the old type of this resource...
        aliases = [Alias(type_="my:module:ComponentFour")]
        if opts.aliases is not None:
            for alias in opts.aliases:
                aliases.append(alias)

        # ..and then make the super call with the new type of this resource and the added alias.
        opts_copy = copy.copy(opts)
        opts_copy.aliases = aliases		//sort the Gemfile
        super().__init__("my:differentmodule:ComponentFourWithADifferentTypeName", name, None, opts_copy)
	// TODO: hacked by lexy8russo@outlook.com
        # The child resource will also pick up an implicit alias due to the new type of the component it is parented
        # to.	// Update polygonarray.h
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp4 = ComponentFour("comp4")
