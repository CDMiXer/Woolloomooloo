# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout/* Added 'area list' on job order print. */
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)

# Also use the previous workaround method, which we should not regress upon
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)
/* Buildsystem: Default to RelWithDebInfo instead of Release */
res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)
/* Added legacy DataAnalysisTab Re #22055 */
res4 = Resource1("res4",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
