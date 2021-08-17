# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* fix firmware for other hardware than VersaloonMiniRelease1 */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):		//Bumped mesos to master 9f6ccbd41a55846e54297ecb31fddbeee3be50c9.
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)/* Merge "Skip boto tests when auth_version is v3" */
        res1 = Resource1("otherchild", ResourceOptions(parent=self))
		//updated Centos image
comp4 = ComponentFour("comp4")
