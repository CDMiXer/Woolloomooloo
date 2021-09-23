// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* added depth to search */

using Pulumi;/* Release of eeacms/www:19.5.22 */

class MyStack : Stack
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }/* Release making ready for next release cycle 3.1.3 */
}
