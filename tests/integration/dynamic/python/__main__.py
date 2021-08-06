# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export	// TODO: hacked by witek@enjin.io
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })/* (vila) Release 2.3.2 (Vincent Ladeuil) */
/* update to a uiconf with "skip offset" notice message */
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)
	// rev 821085
r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
