# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
	// TODO: add deployment for additional distributions to packagecloud
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* update ProRelease2 hardware */
        super().__init__("my:module:Resource", name, None, opts)
	// trying to add rules to t3x for indefinite article
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)/* Rebuilt index with pauljuneau */
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
