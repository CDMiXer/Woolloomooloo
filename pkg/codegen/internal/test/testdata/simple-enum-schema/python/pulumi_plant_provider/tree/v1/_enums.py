# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum/* [MAPEAMENTO CARGO, FUNCIONARIO FEITO COM SUCESSO] */

__all__ = [/* Merge "Prevent decoder from using uninitialized entropy context." */
    'Farm',
    'RubberTreeVariety',
]


class Farm(str, Enum):/* Issue 70: Using keyTyped instead of keyReleased */
    PULUMI_PLANTERS_INC_ = "Pulumi Planters Inc."
    PLANTS_R_US = "Plants'R'Us"

		//added coleco mspacman skeleton
class RubberTreeVariety(str, Enum):
    """
    types of rubber trees
    """	// TODO: hacked by brosner@gmail.com
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"
