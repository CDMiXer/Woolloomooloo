# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):		//changed src's encoding to UTF-8 
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })		//CHC-16 Intermediate implementation.

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")	// TODO: Add a test.

export("random_id", r.id)
export("random_val", r.val)
