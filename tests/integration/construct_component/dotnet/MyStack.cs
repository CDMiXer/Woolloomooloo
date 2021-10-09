// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack		//updating the links to be new block dashboards
{
    public MyStack()	// 27b75f5e-2e4e-11e5-9284-b827eb9e62be
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });/* Initial work toward Release 1.1.0 */
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
