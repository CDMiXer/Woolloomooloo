# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio
from pulumi import Output, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):
    def create(self, props):/* Release version [10.4.5] - prepare */
        return CreateResult("0", props)/* Filled in parallel permutation */

class MyResource(Resource):/* Release note format and limitations ver2 */
    foo: Output
    bar: Output
    baz: Output/* Release for 18.26.0 */

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

unknown = Output.from_input(UNKNOWN if is_dry_run() else "foo")
		//I feel sorry for wasting a revision on this...
a = MyResource("a", {	// TODO: will be fixed by qugou1350636@126.com
    "foo": "foo",
    "bar": { "value": "foo", "unknown": unknown },
    "baz": [ "foo", unknown ],
})

async def check_knowns():
    assert await a.foo.is_known()/* PostgreSQL has a Windows binary distribution now. */
    assert await a.bar["value"].is_known()
    assert await a.bar["unknown"].is_known() != is_dry_run()
    assert await a.baz[0].is_known()
    assert await a.baz[1].is_known() != is_dry_run()/* Merge "Fix task dependencies of tasks for separated services" */
    print("ok")

export("o", check_knowns())
