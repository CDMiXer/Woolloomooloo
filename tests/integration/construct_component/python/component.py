# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]/* Released OpenCodecs 0.84.17325 */

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):	// TODO: 631f7b34-2e68-11e5-9284-b827eb9e62be
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)		//Created tests for file request
