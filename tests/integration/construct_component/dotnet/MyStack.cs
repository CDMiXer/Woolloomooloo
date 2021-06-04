// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack/* New Version 1.3 Released! */
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });	// TODO: will be fixed by boringland@protonmail.ch
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });	// Adds LICENSE.md.
    }
}
