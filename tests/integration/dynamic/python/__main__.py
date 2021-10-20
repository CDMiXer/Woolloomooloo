# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult		//Update README, expected features

class RandomResourceProvider(ResourceProvider):/* Update to version 1.0 for First Release */
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str		//now with proper c# highlighting
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)
	// TODO: [release 0.28.0-RC1] update timestamp and build numbers [ci skip]
r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
