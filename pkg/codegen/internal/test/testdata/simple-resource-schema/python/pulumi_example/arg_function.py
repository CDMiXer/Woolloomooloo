# coding=utf-8
# *** WARNING: this file was generated by test. ***		//3ea009c4-2e4d-11e5-9284-b827eb9e62be
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime		//Merge "Update Trove Installation guide"
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import Resource/* bugfix: client sock timeout error */

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]

@pulumi.output_type
class ArgFunctionResult:		//Try Java 16
    def __init__(__self__, result=None):
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")/* Beta Release */
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Optional['Resource']:
        return pulumi.get(self, "result")/* Updated README to include copyright information */
	// Transfer controller updated

class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(	// TODO: will be fixed by joshua@yottadb.com
            result=self.result)


def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['arg1'] = arg1
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value	// TODO: REQUEST FIX PIM NO 59

    return AwaitableArgFunctionResult(
        result=__ret__.result)
