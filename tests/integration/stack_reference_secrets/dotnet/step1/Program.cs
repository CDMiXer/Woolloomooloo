// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Make sure the --mail option gets passed to the controller's build method. */

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
/* [MERGE] tools.convert: allow to use relativedelta in XML files. */
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
