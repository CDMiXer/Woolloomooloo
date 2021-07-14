// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// added comment //output
using Pulumi;
	// TODO: Delete Help-GUI V3.exe
class MyStack : Stack
{/* Release 0.95.148: few bug fixes. */
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });		//7cd075d9-2e4f-11e5-98c9-28cfe91dbc4b
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
