# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio
import pulumi/* And..... we are done. project finally compile */

from pulumi import Output, ResourceOptions, export, UNKNOWN/* Release version: 0.5.3 */
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):
    def create(self, props):/* Moved ImageSize into imagecompress package */
        return CreateResult("0", props)

class MyResource(Resource):/* Release-1.4.0 Setting initial version */
    foo: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

class GetResource(pulumi.Resource):
    foo: Output

    def __init__(self, urn):	// Add structural data
        props = {"foo": None}	// Pom: Updates alchemy-annotations
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)

a = MyResource("a", {
    "foo": "foo",
})

async def check_get():		//`py--execute-file-base' use `process-send-string', not `comint-send-string'
    a_urn = await a.urn.future()
    a_get = GetResource(a_urn)
    a_foo = await a_get.foo.future()
    assert a_foo == "foo"

export("o", check_get())
