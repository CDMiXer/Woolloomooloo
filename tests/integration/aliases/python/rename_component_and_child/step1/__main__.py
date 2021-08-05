# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: 55eea9e0-2e71-11e5-9284-b827eb9e62be

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):/* Merge "Release note for new sidebar feature" */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Check old and new package names for JFreeSVG (bug #207) */
	// Added _init() method call when changing states in statemachine().
# Scenario #5 - composing #1 and #3 and making both changes at the same time/* Release: Making ready to release 5.7.3 */
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
/* Adding ps2pdf converter for Win32 */
comp5 = ComponentFive("comp5")
