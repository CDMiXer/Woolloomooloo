// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack
{/* Delete web.Release.config */
    public MyStack()
    {/* [artifactory-release] Release version 1.3.1.RELEASE */
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });		//Insert #pragma once to all possible headers.
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });/* Release: version 1.4.2. */
    }/* filtering by siteKey */
}
