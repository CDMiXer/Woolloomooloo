# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi/* extension not persisted */
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union		//Delete wN.png
from . import _utilities, _tables
from . import Resource

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]
	// TODO: [uk] simple replace rule improvements
@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")
        pulumi.set(__self__, "result", result)
	// 3f0062de-2e6a-11e5-9284-b827eb9e62be
    @property
    @pulumi.getter		//Merge "JSDuck-ify /resources/mediawiki.language/*"
    def result(self) -> Optional['Resource']:
        return pulumi.get(self, "result")


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(	// TODO: hacked by witek@enjin.io
            result=self.result)


def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:	// TODO: hacked by sbrichards@gmail.com
    """/* You can title levels */
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['arg1'] = arg1	// Update space.py
    if opts is None:		//Delete CoordinateInputDialog$1.class
        opts = pulumi.InvokeOptions()	// TODO: Удалена лиценция MIT
    if opts.version is None:
        opts.version = _utilities.get_version()		//Rename Apply -> DirectApply
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        result=__ret__.result)
