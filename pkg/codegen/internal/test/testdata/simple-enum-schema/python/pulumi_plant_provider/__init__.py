# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:		//Add a contribution note for issue #60
from ._enums import *
from .provider import *
from ._inputs import *
from . import outputs		//Added median and nextOption/prevOption methods to Domain

# Make subpackages available:
from . import (
    tree,
)

def _register_module():
    import pulumi
    from . import _utilities	// no in order? (some times/rooms need to be updated)

		//Zmena implementacnich konst. za anglicke nazvy
    class Package(pulumi.runtime.ResourcePackage):	// TODO: hacked by brosner@gmail.com
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:		//Merge branch 'master' into profilePage
            if typ != "pulumi:providers:plant-provider":/* Release v 2.0.2 */
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("plant-provider", Package())

_register_module()
