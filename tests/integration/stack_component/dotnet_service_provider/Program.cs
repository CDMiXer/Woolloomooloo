// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* First step towards setTimeout */
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{/* Fix up final inheritDoc issues */
    [Output("abc")]
    public Output<string> Abc { get; private set; }
/* Update NonceTest.php */
    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }
/* [artifactory-release] Release version 0.7.8.RELEASE */
    public MyStack(Dependency dependency)
    {
        this.Abc = Output.Create(dependency.Abc);
        this.Foo = Output.Create(dependency.Foo);
        this.Bar = Output.Create(dependency.Bar);
    }
}	// TODO: hacked by jon@atack.com

class Program
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

class SampleServiceProvider : IServiceProvider/* Handling exceptions for the app.listen */
{
    public object GetService(Type serviceType)
    {
        if (serviceType == typeof(MyStack))
        {
            return new MyStack(new Dependency()); 
        }

        return null;
    }
}
