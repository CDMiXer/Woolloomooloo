# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio/* - Release 1.6 */
from pulumi import Output, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult/* New rc 2.5.10~rc7  (set master table to 14) */
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):/* Updated C# Examples for Release 3.2.0 */
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):/* Release areca-7.2.5 */
    foo: Output
    bar: Output
    baz: Output

    def __init__(self, name, props, opts = None):/* reverted targetSDK */
        super().__init__(MyProvider(), name, props, opts)

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")

a = MyResource("a", {
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],
})
/* Inschrijf button */
async def check_knowns():
    assert await a.foo.is_known()
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()/* Create userAPI.h */
    print("ok")

export("o", check_knowns())
