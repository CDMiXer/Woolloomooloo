# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export/* Release for v40.0.0. */
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):/* Release of eeacms/ims-frontend:0.6.1 */
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")		//1. Functionalities on Inventory tab
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str/* Tests improved. Status 401 removed. */
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)/* Update richards.py */

r = Random("foo")/* New Release (0.9.9) */

export("random_id", r.id)/* Fix class method syntax that caused MSVC compiler errors. */
export("random_val", r.val)
