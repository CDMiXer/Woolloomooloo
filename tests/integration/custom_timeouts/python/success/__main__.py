# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)	// Update Apply-WindowsUpdates.ps1

# Attempt to create a resource with a CustomTimeout
res1 = Resource1("res1",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(create='30m'))		//Merge "Use same MANAGER_TOPIC variable"
)

# Also use the previous workaround method, which we should not regress upon
res2 = Resource1("res2",
    opts=ResourceOptions(custom_timeouts={'create': '15m', 'delete': '15m'})
)
	// TODO: Merge branch 'master' of https://shineangelic@code.google.com/p/souliss/
res3 = Resource1("res3",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(update='30m'))
)

res4 = Resource1("res4",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))/* add comment to the code */
)
