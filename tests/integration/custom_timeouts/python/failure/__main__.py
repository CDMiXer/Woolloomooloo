# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: 4.2.5 no-starve solution
from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):	// Give some Batl examples and comparison
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* Release 1.0.66 */

# Attempt to create a resource with a CustomTimeout that should fail	// Implemented archive.org search performer.
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')
)
