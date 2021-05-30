// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: Merge "#1282 Prevention Updates" into RELEASE_15_BETA

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;/* updating toolbox */

class MyStack : Stack
{
    [Output("abc")]		//Added backend and frontend filters
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }
	// TODO: 2940db94-2e48-11e5-9284-b827eb9e62be
    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();/* stm32 radio:add touch.c */
}
