# coding=utf-8	// TODO: will be fixed by why@ipfs.io
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***
		//Updated LAS2peer version
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables		//readme 1.2
from . import Resource

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]
		//Update 116. Populating Next Right Pointers in Each Node
@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")/* Release version to store */
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Optional['Resource']:
        return pulumi.get(self, "result")


class AwaitableArgFunctionResult(ArgFunctionResult):		//Fire off an index change delta for files that get committed
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(
            result=self.result)


def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """
    Use this data source to access information about an existing resource.	// TODO: Render resource loading progress
    """
    __args__ = dict()
    __args__['arg1'] = arg1
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:/* Target i386 and Release on mac */
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        result=__ret__.result)		//0a4b82b4-2e6d-11e5-9284-b827eb9e62be
