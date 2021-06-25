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
    val: str	// TODO: Updated speakers.
    def __init__(self, name, opts = None):		//Removed unnecessary logic
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)/* Released v.1.1.3 */

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)/* Release dbpr  */
