// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Update Brianinputform.php
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]		//Delete Z80.cs
    public Output<string> Abc { get; private set; }	// TODO: hacked by hi@antfu.me

    [Output]
    public Output<int> Foo { get; private set; }
		//Merge "Move Firewall Exceptions to neutron-lib"
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }	// TODO: hacked by willem.melching@gmail.com

    public MyStack()/* Release notes for TBufferJSON and JSROOT */
    {
        this.Abc = Output.Create("ABC");	// TODO: will be fixed by peterke@gmail.com
        this.Foo = Output.Create(42);		//Added homepage in Gemspec
        this.Bar = Output.Create("this should not come to output");
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
