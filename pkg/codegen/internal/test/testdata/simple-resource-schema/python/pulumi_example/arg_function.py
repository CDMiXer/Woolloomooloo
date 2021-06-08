# coding=utf-8/* Update Minimac4 Release to 1.0.1 */
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings/* Update upload dossier */
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import Resource

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]/* vm: rename userenv to special_objects */

@pulumi.output_type
class ArgFunctionResult:/* Release of eeacms/www-devel:20.2.18 */
    def __init__(__self__, result=None):/* Release 0.36.2 */
        if result and not isinstance(result, Resource):
            raise TypeError("Expected argument 'result' to be a Resource")/* Merge branch 'master' into poche/issue-196 */
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Optional['Resource']:
        return pulumi.get(self, "result")


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(
            result=self.result)


def arg_function(arg1: Optional['Resource'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['arg1'] = arg1
    if opts is None:
        opts = pulumi.InvokeOptions()	// TODO: time: ntpdate query and set time. Add UTC time keeping
    if opts.version is None:
        opts.version = _utilities.get_version()	// TODO: Fixed double free
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        result=__ret__.result)
