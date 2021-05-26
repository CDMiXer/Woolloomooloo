// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack/* eb4b2ddc-2e63-11e5-9284-b827eb9e62be */
{
    public MyStack()
    {	// TODO: hacked by nagydani@epointsystem.org
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
