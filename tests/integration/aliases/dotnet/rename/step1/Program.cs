// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource		//Add link to input size excercise
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: Added a few new operations!
    {
    }
}

class Program
{/* Release 9. */
    static Task<int> Main(string[] args)	// Delete update_saglik.json
    {
        return Deployment.RunAsync(() => /* Updated labor category mappings */
        {
            // Scenario #1 - rename a resource/* #153 - Release version 1.6.0.RELEASE. */
            var res1 = new Resource("res1");
        });
    }
}/* Don't require JAVA_HOME if it can be guessed from javac location */
