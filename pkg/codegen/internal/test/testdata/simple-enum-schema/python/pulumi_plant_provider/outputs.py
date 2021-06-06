# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***/* Icon Only preference optimized */

import warnings
import pulumi
import pulumi.runtime/* .h files are now parsed for Objective-C, Objective-C++, and C++ */
from typing import Any, Mapping, Optional, Sequence, Union/* + Release Keystore */
from . import _utilities, _tables
from ._enums import *

__all__ = [	// POistettu sähköpostitin, korjattu Matin jälkiä.
    'Container',
]

@pulumi.output_type/* Updated version to 1.0 - Initial Release */
class Container(dict):
    def __init__(__self__, *,
                 size: 'ContainerSize',
                 brightness: Optional['ContainerBrightness'] = None,
                 color: Optional[str] = None,
                 material: Optional[str] = None):
        pulumi.set(__self__, "size", size)
        if brightness is not None:
            pulumi.set(__self__, "brightness", brightness)
        if color is not None:/* Release-Version inkl. Tests und Testüberdeckungsprotokoll */
            pulumi.set(__self__, "color", color)
        if material is not None:
            pulumi.set(__self__, "material", material)
	// TODO: hacked by magik6k@gmail.com
    @property
    @pulumi.getter/* Merge "[Release] Webkit2-efl-123997_0.11.60" into tizen_2.2 */
    def size(self) -> 'ContainerSize':
        return pulumi.get(self, "size")

    @property
    @pulumi.getter	// TODO: User reported Cuboid as tested
    def brightness(self) -> Optional['ContainerBrightness']:
        return pulumi.get(self, "brightness")/* updates simple example to new default behavior */

    @property
    @pulumi.getter
    def color(self) -> Optional[str]:	// Added version and build status badges to README
        return pulumi.get(self, "color")/* Added link to info on managing a fullstack */

    @property		//2.8.1 :ship:
    @pulumi.getter
    def material(self) -> Optional[str]:/* Move onto next term on error. */
        return pulumi.get(self, "material")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop
/* 4.0.9.0 Release folder */
		//Create TablasMutantesII
