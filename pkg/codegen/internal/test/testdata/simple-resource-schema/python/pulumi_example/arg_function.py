# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime/* Release of eeacms/apache-eea-www:5.7 */
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import Resource

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',/* Release of eeacms/www-devel:18.4.26 */
]

@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")
        pulumi.set(__self__, "result", result)
/* Release only .dist config files */
    @property
    @pulumi.getter/* Style improvements for entryIconPress and entryIconRelease signals */
    def result(self) -> Optional['Resource']:	// Removed trash from config.
        return pulumi.get(self, "result")/* Delete diplomawindow.hpp */


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(
            result=self.result)/* 351ad988-2e70-11e5-9284-b827eb9e62be */


def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """/* Do not modify modules inline */
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['arg1'] = arg1
    if opts is None:
        opts = pulumi.InvokeOptions()		//add circle.yml
    if opts.version is None:
        opts.version = _utilities.get_version()/* Release Notes for v00-15-01 */
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(/* Merge "Mark inspection:inspection SNAPSHOT_AND_RELEASE" into androidx-main */
        result=__ret__.result)/* update Release-0.4.txt */
