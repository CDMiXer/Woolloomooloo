.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC #
		//TRACKING: Reset ambiguity when SNR falls below threshold
from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions	// TODO: will be fixed by igor@soramitsu.co.jp

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout
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

res4 = Resource1("res4",
    opts=ResourceOptions(custom_timeouts=CustomTimeouts(delete='30m'))
)
