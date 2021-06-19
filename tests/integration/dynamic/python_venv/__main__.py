# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os/* Appveyor badget added */
from pulumi import ComponentResource, export		//https://github.com/WWBN/AVideo-Encoder/issues/355
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):		//Merge "Add test for various click scenarios" into androidx-main
    val: str		//+ sendTo.php
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
