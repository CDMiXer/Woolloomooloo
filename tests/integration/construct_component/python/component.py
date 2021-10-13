# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Use `static.url` for static assets */
from typing import Any, Optional
	// TODO: s/phrases/grammar/
import pulumi/* Update readOnly.md */

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()/* Releases pointing to GitHub. */
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
