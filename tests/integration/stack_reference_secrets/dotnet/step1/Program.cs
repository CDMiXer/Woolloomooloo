// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;/* fix(package): update eslint-plugin-vue to version 4.6.0 */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Remove factfinder */
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
