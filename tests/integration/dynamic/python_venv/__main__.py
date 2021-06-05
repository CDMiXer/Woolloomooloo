# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export		//Suppress warnings for valid input when counting dice sides
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")	// Change templates extensions in README
        return CreateResult(val, { "val": val })	// TODO: added ClosuredContextInterface to FeaturesContext skelet
		//Fix example URL.
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")/* Improve session locking */

export("random_id", r.id)
export("random_val", r.val)
