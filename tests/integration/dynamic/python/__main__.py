# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release of eeacms/forests-frontend:1.7-beta.2 */

import binascii
import os
from pulumi import ComponentResource, export	// filter() really belonged in CollectionUtil; can take any object type
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })
/* Release areca-7.0.9 */
class Random(Resource):
    val: str	// dd68b58a-2e9b-11e5-b79d-a45e60cdfd11
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
