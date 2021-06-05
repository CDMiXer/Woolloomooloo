// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }	// TODO: emails: clarify the `notify_new_draft` action

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }/* Folder structure of biojava4 project adjusted to requirements of ReleaseManager. */

    public MyStack()
    {
        this.Abc = Output.Create("ABC");/* Create org.eclipse.core.resources.prefs */
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }/* Split LightWindow into DecoratedWindow (unthemed), LightWindow and DarkWindow */
}

class Program
{/* send X-Ubuntu-Release to the store */
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();/* Initial commit of period add / diff. Still needs to throw exceptions. */
}
