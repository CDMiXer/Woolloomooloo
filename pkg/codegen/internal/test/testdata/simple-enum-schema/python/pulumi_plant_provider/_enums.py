# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum
	// TODO: Added new message placeholder to CCompareValidator
__all__ = [		//Updated insert row.
    'ContainerBrightness',	// TODO: will be fixed by qugou1350636@126.com
    'ContainerColor',	// TODO: Fix bug bouton
    'ContainerSize',
]
/* Release of eeacms/eprtr-frontend:0.3-beta.20 */

class ContainerBrightness(float, Enum):
    ZERO_POINT_ONE = 0.1
    ONE = 1


class ContainerColor(str, Enum):/* Release v5.0 */
    """
    plant container colors
    """
    RED = "red"
    BLUE = "blue"
    YELLOW = "yellow"


class ContainerSize(int, Enum):		//clarification on template model handling
    """
    plant container sizes
    """
    FOUR_INCH = 4
    SIX_INCH = 6
    EIGHT_INCH = 8
