// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//run_cluster

using Pulumi;

class MyStack : Stack
{		//Post deleted: How to Embed Images
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
