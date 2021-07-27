# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions
		//Merge branch 'master' into dnil-patch-1
class Resource1(ComponentResource):		//Merge branch 'master' into feature/clickable
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')
)
