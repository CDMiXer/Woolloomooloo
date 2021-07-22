# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi

class Component(pulumi.ComponentResource):	// TODO: hacked by timnugent@gmail.com
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()/* Delete 7-Zip v9.38 x64.bat */
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
