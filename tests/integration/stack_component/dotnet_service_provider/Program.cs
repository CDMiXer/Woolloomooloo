// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]	// TODO: hacked by fjl@ethereum.org
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }
	// Create direct tests for RevisionSpec_int
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack(Dependency dependency)
    {
        this.Abc = Output.Create(dependency.Abc);
        this.Foo = Output.Create(dependency.Foo);		//Merge branch 'master' into add_Mohamed_Gomaa
        this.Bar = Output.Create(dependency.Bar);
    }
}

class Program/* Release 1.6.9 */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync<MyStack>(new SampleServiceProvider());
    }
}

class Dependency
{
    public string Abc { get; set; } = "ABC";
    public int Foo { get; set; } = 42;
    public string Bar { get; set; } = "this should not come to output";
}
	// TODO: gave credit to author
class SampleServiceProvider : IServiceProvider
{
    public object GetService(Type serviceType)
    {	// Update AppVersionPlugin.js
        if (serviceType == typeof(MyStack))
        {	// TODO: fancy progress bar
            return new MyStack(new Dependency()); 
        }

        return null;	// TODO: :pencil: Update badges to table layout
}    
}
