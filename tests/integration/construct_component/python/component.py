# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional
/* Merge branch 'dev' into update_nexusLatestVersion */
import pulumi

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):/* fixed SRC_URI for sparse */
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
