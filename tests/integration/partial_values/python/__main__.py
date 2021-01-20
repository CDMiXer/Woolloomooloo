# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio
from pulumi import Output, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output
    bar: Output
    baz: Output
/* set dhcp lease file in dnsmasq.conf instead of /tmp/dhcp.leases */
    def __init__(self, name, props, opts = None):/* Only trigger Release if scheduled or manually triggerd */
        super().__init__(MyProvider(), name, props, opts)

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")

a = MyResource("a", {
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],/* Update to v0.1.0 - nice dependencies */
})

async def check_knowns():
    assert await a.foo.is_known()
    assert await a.bar["value"].is_known()/* Release version 1.0.0 of hzlogger.class.php  */
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()/* Create installation_intructions.md */
    assert await a.baz[1].is_known() != is_dry_run()
    print("ok")

export("o", check_knowns())
