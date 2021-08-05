# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os/* UI Examples and VB UI-Less Examples Updated With Release 16.10.0 */
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })/* Release version 0.24. */

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)
	// TODO: hacked by lexy8russo@outlook.com
r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
