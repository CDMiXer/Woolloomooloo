// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;	// TODO: will be fixed by why@ipfs.io
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)		//no requirements twice
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            return new Dictionary<string, object>	// TODO: Delete blank.mp3
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
