# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 0.8.5 Release for Custodian (#54) */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE
/* Release: update to 4.2.1-shared */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time/* Modified the Deadline so it handles non 0 origin and complements Release */
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
