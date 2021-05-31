# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// improve device model spec

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// TODO: will be fixed by fjl@ethereum.org

# Attempt to create a resource with a CustomTimeout		//Proximal child/sibling inherits definition status from focus concept.
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)

# Also use the previous workaround method, which we should not regress upon
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)
	// Merge "Simplify contentToString function"
res4 = Resource1("res4",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))	// Update Readme.md and see if I'm using git properly.
)
