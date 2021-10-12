// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;	// TODO: Imported Upstream version 0.6.0.1

class Program/* Release LastaTaglib-0.7.0 */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* README: Add the GitHub Releases badge */
        {
            return new Dictionary<string, object>
            {
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
