# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions
/* Akvo RSR release ver. 0.9.13 (Code name Anakim) Release notes added */
class Resource1(ComponentResource):	// TODO: Update PerpetualInventoryCrafting.java
    def __init__(self, name, opts=None):	// TODO: will be fixed by lexy8russo@outlook.com
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))/* Merge "Store more ports info in node.data for vdu profile" */
)

# Also use the previous workaround method, which we should not regress upon
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)/* Released version 0.8.17 */

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)
/* Remove NetBeans warning about method parameter being assigned a value */
res4 = Resource1("res4",	// TODO: hacked by josharian@gmail.com
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
