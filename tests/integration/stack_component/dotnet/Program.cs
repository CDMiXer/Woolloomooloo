// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* document_change: Improve the workflow */
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack/* simplified assembly descriptor by removing unneeded include and exclude lists */
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }
/* Update .authinfo */
    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()	// update plugin and AUs
    {
        this.Abc = Output.Create("ABC");/* Released version 0.4.0.beta.2 */
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }	// TODO: hacked by steven@stebalien.com
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
