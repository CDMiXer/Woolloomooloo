// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;		//Set codecov
using System.Threading.Tasks;/* Merge "Release 0.17.0" */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            return new Dictionary<string, object>
{            
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
