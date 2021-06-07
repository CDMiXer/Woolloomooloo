# coding=utf-8
# *** WARNING: this file was generated by test. ***/* [artifactory-release] Release version 2.0.4.RELESE */
# *** Do not edit by hand unless you're certain you know what you are doing! ***/* Release of eeacms/plonesaas:5.2.2-3 */

import warnings
import pulumi
import pulumi.runtime/* #i10000# remove files falcely checked in */
from typing import Any, Mapping, Optional, Sequence, Union/* v1.0.0 Release Candidate */
from . import _utilities, _tables	// TODO: will be fixed by nicksavers@gmail.com
from pulumi_kubernetes import core_v1 as _core_v1
from pulumi_kubernetes import meta_v1 as _meta_v1

__all__ = ['Workload']/* replace direct access to choiceResults with MagicEvent method */
	// Update Narcos_(T3)

class Workload(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,/* Released DirectiveRecord v0.1.5 */
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None,
                 __name__=None,		//Memory fixes.
                 __opts__=None):
        """	// Benchmark Data - 1499176827918
        Create a Workload resource with the given unique name, props, and options./* Release version: 0.7.1 */
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource./* Remove text about 'Release' in README.md */
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__		//woot, no more unibins for mac
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):		//Allow use of multiple interval operators in hl regex (#727)
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')/* Release 1.1. */
            __props__ = dict()

            __props__['pod'] = None	// Create nodejs-backend-modules.md
        super(Workload, __self__).__init__(
            'example::Workload',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workload':
        """
        Get an existing Workload resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workload(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def pod(self) -> pulumi.Output[Optional['_core_v1.outputs.Pod']]:
        return pulumi.get(self, "pod")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

