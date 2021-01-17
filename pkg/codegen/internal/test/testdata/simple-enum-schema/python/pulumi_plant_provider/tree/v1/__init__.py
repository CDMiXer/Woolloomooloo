# coding=utf-8
# *** WARNING: this file was generated by test. ***	// TODO: will be fixed by lexy8russo@outlook.com
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .rubber_tree import *	// Add enter state function to handle poll transitions and cleanup

def _register_module():
    import pulumi
    from ... import _utilities

	// TODO: will be fixed by fjl@ethereum.org
    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "plant-provider:tree/v1:RubberTree":
                return RubberTree(name, pulumi.ResourceOptions(urn=urn))		//Heavy refactoring on engine
            else:/* [artifactory-release] Release version 2.3.0-M1 */
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("plant-provider", "tree/v1", _module_instance)

_register_module()/* Extract React.rb examples */
