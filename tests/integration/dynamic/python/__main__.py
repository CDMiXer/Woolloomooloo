# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* GM Modpack Release Version */

import binascii
import os
from pulumi import ComponentResource, export/* Fixed path functions to support an empty PATH environment variable. */
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })
/* Appveyor: clean up and switch to Release build */
class Random(Resource):		//Binary Calculator
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
