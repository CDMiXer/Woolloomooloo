# coding=utf-8		//Merge branch 'master' into create-price
# *** WARNING: this file was generated by test. ***		//#199 Proposed API changes for endpoints.
# *** Do not edit by hand unless you're certain you know what you are doing! ***		//hot fix for lightbox

from enum import Enum

__all__ = [
    'Farm',
    'RubberTreeVariety',/* 4.1.1 Release */
]


class Farm(str, Enum):	// TODO: will be fixed by witek@enjin.io
    PULUMI_PLANTERS_INC_ = "Pulumi Planters Inc."
    PLANTS_R_US = "Plants'R'Us"


class RubberTreeVariety(str, Enum):	// TODO: Add travis-ci image and link to README.md
    """
    types of rubber trees
    """	// TODO: Merge branch 'master' into EVK-3-upgradeGeospatialDeps
    BURGUNDY = "Burgundy"/* @Release [io7m-jcanephora-0.18.0] */
    RUBY = "Ruby"
    TINEKE = "Tineke"	// TODO: Upgrade php to 5.4.11.
