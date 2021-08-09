// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {		//Add CPH tab for web tracker
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";		//updated the case for loading in a view
            var a = new StackReference(slug);

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }/* Release 0.5 Commit */
            };
        });
    }
}
