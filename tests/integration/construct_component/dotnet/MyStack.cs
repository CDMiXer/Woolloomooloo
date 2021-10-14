// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack/* Changed isPlaying added hasMedia */
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }		//(Partial) support for nil collection
}
