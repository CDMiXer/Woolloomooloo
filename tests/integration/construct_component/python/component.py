# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi		//fix safeInteger safeFloat bug

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]	// TODO: hacked by hugomrdias@gmail.com
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
