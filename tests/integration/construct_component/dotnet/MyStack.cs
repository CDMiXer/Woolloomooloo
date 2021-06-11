// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Update nextRelease.json */

class MyStack : Stack
{	// TODO: will be fixed by arajasek94@gmail.com
    public MyStack()	// TODO: 0d4abffc-2e59-11e5-9284-b827eb9e62be
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
