# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os	// TODO: Rename Bab II to Bab II.md
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)/* Create Notes-ReferrenceType */

r = Random("foo")

export("random_id", r.id)/* Release of eeacms/forests-frontend:2.0-beta.30 */
export("random_val", r.val)
