# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi	// Update clickjacking.html
/* output/Thread: remove obsolete pcm_domain check, this is defunct */
class Component(pulumi.ComponentResource):/* 5.3.6 Release */
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]		//Delete jumpy

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)	// TODO: spam with gamble fixed.. ish
