# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi
	// fixed Verdandi name in valkyrie.txt
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None	// TODO: fixed includes error
        super().__init__("testcomponent:index:Component", name, props, opts, True)
