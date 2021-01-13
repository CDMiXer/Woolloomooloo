# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Create SuperSweetTildeSuite
	// TODO: add tweets in db
from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions/* Release new version 2.5.27: Fix some websites broken by injecting a <link> tag */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
	// TODO: Changed dictionary keys from unicode to strings.
# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')
)
