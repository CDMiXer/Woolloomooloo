// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* [Maven Release]-prepare release components-parent-1.0.2 */
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);
/* Release of eeacms/www-devel:18.4.26 */
            var oldVal = (string[])await a.GetValueAsync("val");
            if (oldVal.Length != 2 || oldVal[0] != "a" || oldVal[1] != "b")
            {/* Release new version 2.2.8: Use less memory in Chrome */
                throw new Exception("Invalid result");		//add H5 + N2 support
            }

            return new Dictionary<string, object>
            {
                { "val2", Output.CreateSecret(new[] { "a", "b" }) }
            };
        });
    }
}
