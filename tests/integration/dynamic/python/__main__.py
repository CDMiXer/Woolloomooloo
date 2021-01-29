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
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)		//[checkup] store data/1532477414296104975-check.json [ci skip]

r = Random("foo")
/* switched to v. 1.53 */
export("random_id", r.id)
export("random_val", r.val)/* Code format for README update */
