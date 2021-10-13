// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]	// TWEIAL-264 Update styling of play button.
    public Output<int> Foo { get; private set; }
	// TODO: hacked by seth@sethvargo.com
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }/* Release notes e link pro sistema Interage */
}
	// TODO: will be fixed by nicksavers@gmail.com
class Program	// python : Constant
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
