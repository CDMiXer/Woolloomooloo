# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):	// TODO: Updating build-info/dotnet/buildtools/master for preview2-02521-04
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })
	// TODO: Added TODO item.
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)/* media-libs/freetype: update according portage */

r = Random("foo")

export("random_id", r.id)/* Release 1.1.0-RC2 */
export("random_val", r.val)
