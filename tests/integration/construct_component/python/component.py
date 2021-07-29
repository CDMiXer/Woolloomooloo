# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]		//fixed broken url in index.rst
    childId: pulumi.Output[str]/* Automatic changelog generation for PR #8084 [ci skip] */

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None		//Link to wp.org
        super().__init__("testcomponent:index:Component", name, props, opts, True)
