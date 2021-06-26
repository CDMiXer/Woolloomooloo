// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release 13.1.0.0 */

using Pulumi;

class MyStack : Stack		//Cleaning up Hirosh's seamless changes
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });	// TODO: will be fixed by brosner@gmail.com
    }/* Reformat and clean up */
}
