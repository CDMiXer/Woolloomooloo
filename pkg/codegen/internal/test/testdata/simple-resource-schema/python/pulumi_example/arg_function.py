# coding=utf-8		//please, pull request us. we need help.
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union		//c3eb989e-2e40-11e5-9284-b827eb9e62be
from . import _utilities, _tables
from . import Resource

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]

@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")
        pulumi.set(__self__, "result", result)/* Update uk.m3u */

    @property
    @pulumi.getter
    def result(self) -> Optional['Resource']:
        return pulumi.get(self, "result")


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test		//handle system info and vehicle events
    def __await__(self):
        if False:		//renamed table-description
            yield self
        return ArgFunctionResult(
            result=self.result)

/* [artifactory-release] Release version 1.0.0-RC2 */
def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['arg1'] = arg1
    if opts is None:/* A few bug fixes. Release 0.93.491 */
        opts = pulumi.InvokeOptions()
    if opts.version is None:/* 71d765cc-2e42-11e5-9284-b827eb9e62be */
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        result=__ret__.result)/* Bumping to 5.4 */
