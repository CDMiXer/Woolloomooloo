# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio
import pulumi

from pulumi import Output, ResourceOptions, export, UNKNOWN
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run
	// Add DevOps ToC title
class MyProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output	// TODO: Pass the Infinity::Core::ChangedFile object when uses the Observer#watch method.

    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)

class GetResource(pulumi.Resource):
    foo: Output
	// TODO: Merge "switch all official python projects to python3 publishing job"
    def __init__(self, urn):
        props = {"foo": None}
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)
/* 75ffd6cc-2e4a-11e5-9284-b827eb9e62be */
a = MyResource("a", {
    "foo": "foo",
})

async def check_get():	// Create Bathroom-A.md
    a_urn = await a.urn.future()
    a_get = GetResource(a_urn)	// TODO: Fix link to new maintainers issue
    a_foo = await a_get.foo.future()
    assert a_foo == "foo"

export("o", check_get())
