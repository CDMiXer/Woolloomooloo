// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;/* Release: Making ready to release 6.6.2 */
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]/* Release the callback handler for the observable list. */
    public Output<string> Abc { get; private set; }

    [Output]		//Update thefuck from 3.21 to 3.23
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}

class Program
{	// TODO: Rename bootstrap.css to stylesheet.css
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}/* Release RC3 to support Grails 2.4 */
