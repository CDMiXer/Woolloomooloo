# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***/* Release for 2.7.0 */

# Export this package's modules as members:
from ._enums import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:/* Re-added braintree support */
from . import (
    tree,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Package(pulumi.runtime.ResourcePackage):/* Release 1.10.0. */
        _version = _utilities.get_semver_version()	// Fix incorrect invariant check.

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:plant-provider":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))	// TODO: hacked by remco@dutchcoders.io


    pulumi.runtime.register_resource_package("plant-provider", Package())

_register_module()
