# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout
res1 = Resource1("res1",/* Release 1.7.2: Better compatibility with other programs */
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)
		//Use ordered group by default in CustomContent
# Also use the previous workaround method, which we should not regress upon
res2 = Resource1("res2",		//VLC support
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)

res4 = Resource1("res4",/* Update Release.java */
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)/* Release 3.0.4. */
