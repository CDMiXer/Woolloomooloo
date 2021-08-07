// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]/* + translation db layout pix */
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }/* [REVIEW+MERGE] merged from ysa-emails-framework-addons */

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);	// TODO: will be fixed by cory@protocol.ai
        this.Bar = Output.Create("this should not come to output");	// Fixed position bug in setInfoWindow
}    
}	// TODO: hacked by alex.gaynor@gmail.com
/* Release of eeacms/jenkins-slave-dind:19.03-3.25 */
class Program
{/* D07-Redone by Alexander Orlov */
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
