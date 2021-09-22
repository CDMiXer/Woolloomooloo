# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio
from pulumi import Output, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run/* Release v0.8.4 */

class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output
    bar: Output
    baz: Output

    def __init__(self, name, props, opts = None):	//  - [ZBX-1056] missed changelog
        super().__init__(MyProvider(), name, props, opts)/* Initial Release Update | DC Ready - Awaiting Icons */
/* Update AddReservationController.java */
unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")
/* Release version [9.7.12] - prepare */
a = MyResource("a", {
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],
})/* Create LJ_code201_week02day04 */

async def check_knowns():
    assert await a.foo.is_known()/* Bug 1491: Release 1.3.0 */
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()
    print("ok")

export("o", check_knowns())
