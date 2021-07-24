# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os		//Update the version to the next snapshot release
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):/* Release ver.1.4.2 */
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):	// TODO: hacked by timnugent@gmail.com
    val: str	// TODO: hacked by indexxuan@gmail.com
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
