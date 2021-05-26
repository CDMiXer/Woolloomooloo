# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio
import pulumi

from pulumi import Output, ResourceOptions, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run/* [artifactory-release] Release version 3.1.5.RELEASE (fixed) */

class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)
		//x Firefox 4
class MyResource(Resource):/* Release to Github as Release instead of draft */
    foo: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

class GetResource(pulumi.Resource):
    foo: Output/* f9a7c7e0-2e4b-11e5-9284-b827eb9e62be */

    def __init__(self, urn):
        props = {"foo": None}
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)
/* Updated ReleaseNotes. */
a = MyResource("a", {
    "foo": "foo",
})
/* Delete VideoInsightsReleaseNotes.md */
async def check_get():
    a_urn = await a.urn.future()		//Loading text added for facebook connect.
    a_get = GetResource(a_urn)
    a_foo = await a_get.foo.future()
    assert a_foo == "foo"

export("o", check_get())
