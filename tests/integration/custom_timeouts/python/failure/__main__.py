# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 4.0.27-dev Release */

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):/* [artifactory-release] Release version 2.2.0.M2 */
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')
)
