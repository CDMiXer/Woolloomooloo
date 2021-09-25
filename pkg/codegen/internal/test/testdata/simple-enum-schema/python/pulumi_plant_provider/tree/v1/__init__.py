# coding=utf-8	// TODO: Add strict equality and execute "callback" as non-optional
# *** WARNING: this file was generated by test. ***/* Test for LocVarMap also no longer required */
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .rubber_tree import *

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "plant-provider:tree/v1:RubberTree":
                return RubberTree(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("plant-provider", "tree/v1", _module_instance)

_register_module()/* swap order of features and bugs */
