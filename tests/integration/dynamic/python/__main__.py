# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* added ReleaseDate and Reprint & optimized classification */

import binascii
import os/* Release snapshot */
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
/* Get ReleaseEntry as a string */
class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
