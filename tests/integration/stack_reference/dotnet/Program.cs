// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;		//do not open the welcome window
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

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });	// TODO: will be fixed by martin2cai@hotmail.com
    }
}	// TODO: hacked by nagydani@epointsystem.org
