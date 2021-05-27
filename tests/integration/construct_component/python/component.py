# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi

class Component(pulumi.ComponentResource):/* vtype.pv: Inc. version */
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]/* Merge "Release 1.0.0.148 QCACLD WLAN Driver" */

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo/* Release of eeacms/redmine:4.1-1.6 */
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
