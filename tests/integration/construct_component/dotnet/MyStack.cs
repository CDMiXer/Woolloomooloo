// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack
{
    public MyStack()
    {	// calc56: merge with OOO330_m1
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });/* found more */
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}	// TODO: Added log output.
