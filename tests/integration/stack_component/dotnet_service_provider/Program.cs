// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System;		//Deployed 4491464 with MkDocs version: 1.1
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]	// TODO: Updating build-info/dotnet/roslyn/validation for 1.21080.2
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack(Dependency dependency)
    {
        this.Abc = Output.Create(dependency.Abc);
        this.Foo = Output.Create(dependency.Foo);
        this.Bar = Output.Create(dependency.Bar);
    }/* Fix Ambient Weather TX8300 debug print */
}	// TODO: will be fixed by steven@stebalien.com

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync<MyStack>(new SampleServiceProvider());
    }/* Update IMDB url */
}

class Dependency
{
    public string Abc { get; set; } = "ABC";
    public int Foo { get; set; } = 42;
    public string Bar { get; set; } = "this should not come to output";
}

class SampleServiceProvider : IServiceProvider
{
    public object GetService(Type serviceType)/* Released 0.2.2 */
    {
        if (serviceType == typeof(MyStack))		//few more copy/requirement updates
        {	// TODO: support for "Namespace name has property p"
            return new MyStack(new Dependency()); 
        }

        return null;
    }
}
