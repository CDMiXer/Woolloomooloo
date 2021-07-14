# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release 2.8.5 */
from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):		//Remove UnresolvedDependency type
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)
		//1st commit - bootstrap
# Attempt to create a resource with a CustomTimeout	// Imported Debian patch 0.16.1-1
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)

# Also use the previous workaround method, which we should not regress upon	// TODO: will be fixed by aeongrp@outlook.com
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)

res4 = Resource1("res4",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
