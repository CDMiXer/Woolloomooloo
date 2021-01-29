// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
	// TODO: update review by authenticated user ok
class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }
/* Release v0.0.6 */
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");	// Skin and social link updates
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}	// Added BrowserStack Thanks part to ReadMe

class Program
{		//Increased version number to 1.1.5 and updated dependencies
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}/* Release version 2.2.5.5 */
