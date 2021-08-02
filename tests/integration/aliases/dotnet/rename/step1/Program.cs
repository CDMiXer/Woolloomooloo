// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by cory@protocol.ai
/* 0.9.5 Release */
using System.Threading.Tasks;		//Create Prova.java
using Pulumi;	// moved cookie notice to the bottom of the page

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Release new version to include recent fixes */
    {	// d3cc273c-2e55-11e5-9284-b827eb9e62be
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
