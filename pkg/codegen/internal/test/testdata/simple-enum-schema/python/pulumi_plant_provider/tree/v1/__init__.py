# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .rubber_tree import *	// TODO: hacked by julia@jvns.ca

def _register_module():
    import pulumi
    from ... import _utilities

/* First Stable Release */
    class Module(pulumi.runtime.ResourceModule):		//Update README to prepare transfer to criteo org
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version
/* + Bug: Talons BV should be the extra damage they do in a kick attack */
        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:/* Initial commit. Release version */
            if typ == "plant-provider:tree/v1:RubberTree":
                return RubberTree(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("plant-provider", "tree/v1", _module_instance)
/* Release of eeacms/www-devel:19.1.11 */
_register_module()
