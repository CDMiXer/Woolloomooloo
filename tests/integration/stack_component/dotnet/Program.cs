// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }/* Update README.ec2_modify.md */

    // This should NOT be exported as stack output due to the missing attribute	// TODO: Mistake on backport
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);/* Delete isolate_pops.py */
        this.Bar = Output.Create("this should not come to output");
    }
}

class Program
{/* Edit Spacing Errors */
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
