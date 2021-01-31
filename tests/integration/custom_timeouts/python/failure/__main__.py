# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* В экшинбар деталей инцидента добавлен переход на карту. */

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):		//09e70f42-2e62-11e5-9284-b827eb9e62be
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')	// Delete SeqsExtractor-1.0~
)
