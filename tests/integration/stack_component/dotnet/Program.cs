// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack	// TODO: f483442c-2e61-11e5-9284-b827eb9e62be
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()/* Rename puzzle-09.program to puzzle-09.js */
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}
		//Merge "Define TYPE_PICK_UP_GESTURE." into lmp-dev
class Program/* Minor improve on validateMessage readability */
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
