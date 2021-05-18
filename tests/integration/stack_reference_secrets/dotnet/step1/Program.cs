// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// Update 11.txt
/* Syntax highlight one block. Acknowledge str.format. */
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program		//apache didnt like a url inside a url
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
