// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Add comment to windows doc */

class MyStack : Stack
{
    public MyStack()/* Probably better */
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });	// TODO: hacked by ligi@ligi.de
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });	// TODO: will be fixed by igor@soramitsu.co.jp
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });/* Trying new iteration procedure.... will this work I wonder? */
    }
}/* Create textMe.py */
