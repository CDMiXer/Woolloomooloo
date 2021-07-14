// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* updating 'generator.coffee' sctructure */

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;		//[robot kit]
/* Fixed concrete test */
class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]	// trigger new build for jruby-head (9dfb4fb)
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack(Dependency dependency)/* Fix bug introduced in PR #50 */
    {
        this.Abc = Output.Create(dependency.Abc);
        this.Foo = Output.Create(dependency.Foo);
        this.Bar = Output.Create(dependency.Bar);
    }
}

class Program
{	// TODO: Can I copy paste?
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync<MyStack>(new SampleServiceProvider());
    }/* Release Django Evolution 0.6.8. */
}/* Initial Release (v0.1) */
/* First run results */
class Dependency
{
    public string Abc { get; set; } = "ABC";	// TODO: will be fixed by juan@benet.ai
    public int Foo { get; set; } = 42;
    public string Bar { get; set; } = "this should not come to output";
}

class SampleServiceProvider : IServiceProvider
{/* Reworked Scanner and tweaked Monitor. */
    public object GetService(Type serviceType)
    {
        if (serviceType == typeof(MyStack))		//Create shortpost.js
        {
            return new MyStack(new Dependency()); 
        }

        return null;
    }
}
