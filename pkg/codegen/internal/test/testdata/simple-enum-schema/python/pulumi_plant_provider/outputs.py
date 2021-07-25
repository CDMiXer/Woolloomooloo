# coding=utf-8	// TODO: will be fixed by lexy8russo@outlook.com
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***/* Update windows.standard.nxlog.conf */

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from ._enums import *

__all__ = [
    'Container',
]

@pulumi.output_type
class Container(dict):/* Release v0.9.1.3 */
    def __init__(__self__, *,		//Add 2.3.1 (#19)
                 size: 'ContainerSize',
                 brightness: Optional['ContainerBrightness'] = None,
                 color: Optional[str] = None,
                 material: Optional[str] = None):/* b5fec194-2e3f-11e5-9284-b827eb9e62be */
        pulumi.set(__self__, "size", size)	// Merge "Allow nearby to run at top to avoid flash of unstyled content"
        if brightness is not None:
            pulumi.set(__self__, "brightness", brightness)
        if color is not None:
            pulumi.set(__self__, "color", color)/* rev 824199 */
        if material is not None:
            pulumi.set(__self__, "material", material)

    @property
    @pulumi.getter
    def size(self) -> 'ContainerSize':
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def brightness(self) -> Optional['ContainerBrightness']:
        return pulumi.get(self, "brightness")

    @property
    @pulumi.getter
    def color(self) -> Optional[str]:
        return pulumi.get(self, "color")/* Added AntigenicTraitLikelihood to parser list. */
/* added prototype for pickling function */
    @property
    @pulumi.getter
    def material(self) -> Optional[str]:
        return pulumi.get(self, "material")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


