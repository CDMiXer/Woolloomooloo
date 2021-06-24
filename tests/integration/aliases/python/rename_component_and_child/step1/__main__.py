# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* a32de710-2e68-11e5-9284-b827eb9e62be */

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
	// rename 's' to '_' method @Expression
# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):		//Merge branch 'master' into feature/fix-updateadminprofile-recordtypes
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)	// TODO: Update lambdaJSON.py
        res1 = Resource1("otherchild", ResourceOptions(parent=self))

comp5 = ComponentFive("comp5")
