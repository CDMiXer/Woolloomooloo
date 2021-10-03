// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Increment to 1.5.0 Release */

using System.Collections.Generic;		//Document output_encoding
using System.Threading.Tasks;
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
                { "secret", Output.CreateSecret("secret") },		//On Windows, there is no tzset function, skip on unittest.
            };
        });
    }
}
