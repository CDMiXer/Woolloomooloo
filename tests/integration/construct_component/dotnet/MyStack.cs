// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Added RestingControl (thank you Teemu Salminen) */

class MyStack : Stack
{
    public MyStack()	// Break pane API into sections
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });/* -Corrección de modelos */
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });		//Fix bug [ 1884368 ] festatus command doesn't work on DVB-S (patch supplied)
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
