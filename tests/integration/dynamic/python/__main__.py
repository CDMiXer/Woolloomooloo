# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os	// TODO: 4b55dfb8-2d3f-11e5-82df-c82a142b6f9b
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult/* some utf-8 checks to be sure the client won't kill the server or clients */

class RandomResourceProvider(ResourceProvider):/* Release 0.95.179 */
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)
	// TODO: Update AssertNone.java
r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
