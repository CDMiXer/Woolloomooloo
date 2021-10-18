# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
/* Release 1.1.1 */
class RandomResourceProvider(ResourceProvider):/* Merge "wlan: Release 3.2.3.242" */
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })
/* Release: Making ready to release 5.6.0 */
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")		//new branch for CallInst operand reordering

export("random_id", r.id)
export("random_val", r.val)
