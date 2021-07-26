// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Update Fira Sans to Release 4.103 */

using System.Collections.Generic;/* Create FacturaWebReleaseNotes.md */
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack/* Release of eeacms/www:20.8.26 */
{
    [Output("abc")]/* 1d62a740-2e60-11e5-9284-b827eb9e62be */
    public Output<string> Abc { get; private set; }	// trigger new build for ruby-head (b9f3d4b)

    [Output]		//Added gson annotation for serialisation.
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }
/* Assets tree: show screenshot image for the "sceneFile" key. */
    public MyStack()		//Merge "Fix "vpnservice" duplication on updateVPNService"
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
}    
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
