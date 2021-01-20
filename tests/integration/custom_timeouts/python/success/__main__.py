# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions	// TODO: Function to get MAC addresses & geolocation

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// se cambia nombre tabla

# Attempt to create a resource with a CustomTimeout
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))
)	// Added support for Ubuntu 13.10

# Also use the previous workaround method, which we should not regress upon	// TODO: Started adding the stimulus image capture features to the presenter.
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})	// TODO: Remove readme.md
)

res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)

res4 = Resource1("res4",	// TODO: hacked by why@ipfs.io
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
