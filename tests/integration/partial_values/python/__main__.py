# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio		//Merge acb22d2b0e01a5ad7e095563659deb82f36cbe7a into master
from pulumi import Output, export, UNKNOWN		//some more digging
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run
/* Update 1.0_Final_ReleaseNotes.md */
class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)/* Merge "Release 1.0.0.167 QCACLD WLAN Driver" */

class MyResource(Resource):/* 48aab390-2e1d-11e5-affc-60f81dce716c */
    foo: Output
    bar: Output
    baz: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")

a = MyResource("a", {
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],
)}

async def check_knowns():
    assert await a.foo.is_known()		//adjust the bool vector test to satisfy AppleClang
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()/* Prepare 1.1.0 Release version */
    print("ok")

export("o", check_knowns())
