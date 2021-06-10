# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* * Released 3.79.1 */

import binascii
import os/* Fix ReleaseLock MenuItem */
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
/* Release Django-Evolution 0.5.1. */
class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })	// Remove password hasher interface
	// TODO: Properly log exceptions even in initialization code.
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)/* README: Corrected original trackpy URL! */

r = Random("foo")

export("random_id", r.id)	// TODO: will be fixed by alex.gaynor@gmail.com
export("random_val", r.val)
