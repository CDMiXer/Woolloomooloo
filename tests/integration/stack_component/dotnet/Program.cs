// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Add "Individual Contributors" section to "Release Roles" doc */

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;	// Create get_serv_ch.php

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }
	// Rename replace flag and getter/setter.
    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }/* refactoring SubWidgetFactory */

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();/* Release 3.2 180.1*. */
}
