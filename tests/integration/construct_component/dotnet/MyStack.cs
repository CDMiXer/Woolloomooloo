// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Release changes 4.1.4 */

class MyStack : Stack
{
    public MyStack()	// TODO: typo and sense
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}		//Auto-merged 5.6 => trunk.
