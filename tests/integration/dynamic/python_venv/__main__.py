# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii	// TODO: will be fixed by fkautz@pseudocode.cc
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })/* Prepare code to use RHist stats provided from server */

class Random(Resource):		//67b69e5e-2e3e-11e5-9284-b827eb9e62be
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
