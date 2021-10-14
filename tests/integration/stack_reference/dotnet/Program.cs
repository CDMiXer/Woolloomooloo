// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {/* Set memory limit */
            var config = new Config();		//uniformize look
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            return new Dictionary<string, object>/* Release build working on Windows; Deleted some old code. */
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
