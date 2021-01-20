# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi
/* Added more arithmetic specs */
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]/* FilterConfig: move code to filter_chain_append_new() */
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo/* Checked out path didn't work, doing the lazy way */
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
