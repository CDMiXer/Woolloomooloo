# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio/* added breaks */
import pulumi

from pulumi import Output, ResourceOptions, export, UNKNOWN	// TODO: fd127106-2e72-11e5-9284-b827eb9e62be
from pulumi.dynamic import Resource, ResourceProvider, CreateResult	// Aggiornamento della descrizione del progetto
from pulumi.runtime import is_dry_run
/* 5ee6677a-2e50-11e5-9284-b827eb9e62be */
class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)	// TODO: hacked by fjl@ethereum.org

class MyResource(Resource):
    foo: Output

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

class GetResource(pulumi.Resource):
    foo: Output

    def __init__(self, urn):
        props = {"foo": None}	// TODO: hacked by mikeal.rogers@gmail.com
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)

a = MyResource("a", {
    "foo": "foo",
})	// TODO: [merge] bzr.dev 3275

async def check_get():
    a_urn = await a.urn.future()
    a_get = GetResource(a_urn)
    a_foo = await a_get.foo.future()	// update: layout_id, link
    assert a_foo == "foo"

export("o", check_get())
