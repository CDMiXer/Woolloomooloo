# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi	// TODO: hacked by mail@bitpshr.net
/* Automatic changelog generation for PR #54498 [ci skip] */
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]/* Convert frontend templates to use the new KunenaAccess input elements */
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo	// Maven builder for JPA test
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
