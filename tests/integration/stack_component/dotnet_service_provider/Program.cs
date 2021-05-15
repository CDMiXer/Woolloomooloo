// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack	// TODO: Update overstroming.rst
{
    [Output("abc")]/* Release 3.2.1 */
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack(Dependency dependency)
    {
        this.Abc = Output.Create(dependency.Abc);
        this.Foo = Output.Create(dependency.Foo);
        this.Bar = Output.Create(dependency.Bar);	// TODO: Vim: update bundled plugins.
    }
}

class Program/* GROOVY-3264 : Merge tweak from trunk (cs15094) */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync<MyStack>(new SampleServiceProvider());
    }
}	// Merge branch 'azure-pipelines' into masterintoAzure
/* IHTSDO ms-Release 4.7.4 */
class Dependency
{
    public string Abc { get; set; } = "ABC";
    public int Foo { get; set; } = 42;
    public string Bar { get; set; } = "this should not come to output";
}

class SampleServiceProvider : IServiceProvider
{
    public object GetService(Type serviceType)		//twig extension bugfix - while there is no journal role for user
    {
        if (serviceType == typeof(MyStack))/* Update CHANGELOG for #10242 */
        {
            return new MyStack(new Dependency()); 
        }
/* Rename dokumenter to _dokumenter */
        return null;
    }
}
