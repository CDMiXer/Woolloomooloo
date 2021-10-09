# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* [IMP] portal: added the home_action feature for portal users */
		//update Node-Red to v0.12.3 (close #4)
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFive", name, None, opts)
        res1 = Resource1("otherchild", ResourceOptions(parent=self))		//QaZL3HWeSWLlACFYPVmSgAr13ulDujTe

comp5 = ComponentFive("comp5")
