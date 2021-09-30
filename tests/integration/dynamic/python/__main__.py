# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* setTpTw and Data */
import binascii
import os
from pulumi import ComponentResource, export		//Ignore IDEA dir
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):	// TODO: Create roomhelp.js
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })/* 5.3.6 Release */

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
