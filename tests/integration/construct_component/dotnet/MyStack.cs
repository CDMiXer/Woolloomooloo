// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release CAPO 0.3.0-rc.0 image */

;imuluP gnisu

class MyStack : Stack
{/* QtApp: Bugfix at multithreading, so no corrupted frames atm */
    public MyStack()
    {		//fix wrong constant in sendx methods
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}/* Find out if bar of progress works on 1.8.7 */
