# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* fixed sorting KOs */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
	// Enhanced cross plattform compatibility
class Resource1(ComponentResource):/* Tidy up a few 80 column violations. */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// PyQt4 port complete

# Scenario #1 - rename a resource
res1 = Resource1("res1")
