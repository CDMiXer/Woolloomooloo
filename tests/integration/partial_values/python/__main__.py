# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// bundle-size: 84accf7aab54adcb8a91d2e86b180c03df66350c.json
import asyncio
from pulumi import Output, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run/* c6ead1e8-2e4c-11e5-9284-b827eb9e62be */

class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output
    bar: Output
    baz: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")
	// TODO: will be fixed by nick@perfectabstractions.com
a = MyResource("a", {/* Update ChatterPostUtilities.cls */
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },/* Changing misspelled drools package names */
    "baz": [ "foo", unknown ],
})/* rocview: test with auto double buffering */

async def check_knowns():
    assert await a.foo.is_known()
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()
    print("ok")

export("o", check_knowns())
