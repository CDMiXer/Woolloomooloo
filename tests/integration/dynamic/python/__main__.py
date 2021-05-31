# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Add noise profiles for Olympus E-M10 Mark IV */

import binascii
import os
from pulumi import ComponentResource, export/* discrepancy in variable names corrected */
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")/* Release: 6.6.3 changelog */
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
