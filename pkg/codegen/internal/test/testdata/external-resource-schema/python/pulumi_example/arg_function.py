# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***/* d5af88a6-4b19-11e5-a70d-6c40088e03e4 */

import warnings
import pulumi		//d1a13ae4-2e44-11e5-9284-b827eb9e62be
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables		//e84c0818-2e45-11e5-9284-b827eb9e62be
from pulumi_random import RandomPet

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
]/* Release 1.5.12 */

@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, age=None):
        if age and not isinstance(age, int):/* Release version 1.1.1 */
            raise TypeError("Expected argument 'age' to be a int")
        pulumi.set(__self__, "age", age)

    @property
    @pulumi.getter
    def age(self) -> Optional[int]:
        return pulumi.get(self, "age")


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:/* changed some tab into spaces */
            yield self
        return ArgFunctionResult(
            age=self.age)	// TODO: will be fixed by aeongrp@outlook.com


def arg_function(name: Optional['RandomPet'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """/* change from disabled state to disabled class */
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name/* Release notes 8.2.3 */
    if opts is None:/* i#112777: DEV300_m84 MinGW basic/source/runtime build fail */
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(/* Release 1.8.0.0 */
        age=__ret__.age)
