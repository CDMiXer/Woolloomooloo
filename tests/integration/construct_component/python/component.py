# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi

class Component(pulumi.ComponentResource):/* Remove homepage must not be transparent */
    echo: pulumi.Output[Any]/* Merge "Always export IMAGE_NAME" */
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()	// TODO: Created flipTextSansReverse, it's flipText without the reversing
        props["echo"] = echo
        props["childId"] = None/* Release 1.4.0.8 */
        super().__init__("testcomponent:index:Component", name, props, opts, True)
