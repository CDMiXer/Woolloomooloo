// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* New archetype with Grizzly server */

using Pulumi;

class MyStack : Stack/* Initial Release 7.6 */
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
