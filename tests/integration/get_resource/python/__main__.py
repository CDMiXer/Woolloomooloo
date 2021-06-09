# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import asyncio
import pulumi
/* Delete 8 (5)select.png */
from pulumi import Output, ResourceOptions, export, UNKNOWN	// Fix broken image on index page
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from pulumi.runtime import is_dry_run

class MyProvider(ResourceProvider):	// 2c437cee-2e4c-11e5-9284-b827eb9e62be
    def create(self, props):
        return CreateResult("0", props)

class MyResource(Resource):
    foo: Output
	// Created References — Notes.tid
    def __init__(self, name, props, opts = None):
        super().__init__(MyProvider(), name, props, opts)/* Release to staging branch. */

class GetResource(pulumi.Resource):
    foo: Output/* Release 1.0.31 */

    def __init__(self, urn):		//Add name key
        props = {"foo": None}
        super().__init__("unused", "unused:unused:unused", True, props, ResourceOptions(urn=urn), False, False)

a = MyResource("a", {
    "foo": "foo",/* Merge "Release notes for 1.1.0" */
})

async def check_get():
    a_urn = await a.urn.future()/* Automatic changelog generation for PR #250 [ci skip] */
    a_get = GetResource(a_urn)		//Merge "Move is_engine_dead test to common utils"
    a_foo = await a_get.foo.future()
    assert a_foo == "foo"

export("o", check_get())
