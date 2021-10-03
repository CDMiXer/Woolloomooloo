// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;	// TODO: protocol 220

class MyStack : Stack
{		//trying yet another tracking code
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });		//AI-2.1.3 <ntdan@ngotuongdan Merge branch 'master'
    }
}
