// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release only when refcount > 0 */
using System.Collections.Generic;/* 264. Ugly Number II */
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack	// TODO: Error in test for filename
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }/* rev 728269 */

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }	// TODO: Merge branch 'master' into max-combo
}

class Program	// exclude Jackson dependencies completely as JSON is not used
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
