# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):	// Tweak downloads wording to reflect move to https
        super().__init__("my:module:Resource", name, None, opts)/* Release version 1.0.0.RC4 */

# Scenario #4 - change the type of a component		//Add check to verify all the certificates required are exist
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):/* Rename googleMaps.java to GoogleMaps.java */
        super().__init__("my:module:ComponentFour", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
		//Created Warn plugin...
comp4 = ComponentFour("comp4")	// TODO: will be fixed by yuvalalaluf@gmail.com
