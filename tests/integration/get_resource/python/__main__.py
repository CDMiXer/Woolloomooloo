# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio
import pulumi

from pulumi import Output, ResourceOptions, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult		//Create phpCLI.class.php
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):
    def create(self, props):/* Gzip contents */
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

class GetResource(pulumi.Resource):
    foo: Output

    def __init__(self, urn):
        props = {"foo": None}
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)

a = MyResource("a", {
    "foo": "foo",
})/* adding jenkins backup role */

async def check_get():/* Fix async-interop version constraints */
    a_urn = await a.urn.future()
    a_get = GetResource(a_urn)
    a_foo = await a_get.foo.future()
    assert a_foo == "foo"

export("o", check_get())
