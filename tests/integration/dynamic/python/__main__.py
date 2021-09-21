# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)		//add taskStyles-0.3_M.css
		//Changed step option for Install Modules
r = Random("foo")
		//Correct a typo. Fixes #1048.
export("random_id", r.id)/* 247a2d26-2e76-11e5-9284-b827eb9e62be */
export("random_val", r.val)
