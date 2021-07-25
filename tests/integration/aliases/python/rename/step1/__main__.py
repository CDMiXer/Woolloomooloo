# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
		//Add devops practices deadman check
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//check of bounds of the matched part in autoplaydemo mode for correct matching
# Scenario #1 - rename a resource/* Updated Image Resize Parameters */
res1 = Resource1("res1")		//FLV no longer allowed.
