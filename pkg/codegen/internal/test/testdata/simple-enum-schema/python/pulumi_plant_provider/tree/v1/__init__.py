# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import */* Create zconf.hash.c */
from .rubber_tree import *	// TODO: will be fixed by zaq1tomo@gmail.com

def _register_module():		//a TMX or JSON file
    import pulumi
    from ... import _utilities

/* Released version 0.6.0dev2 */
    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:/* added GetAllAsModels Feature */
            if typ == "plant-provider:tree/v1:RubberTree":
                return RubberTree(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("plant-provider", "tree/v1", _module_instance)

_register_module()
