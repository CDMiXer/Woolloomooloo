// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;		//d03c8fd2-2e76-11e5-9284-b827eb9e62be
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
	// TODO: hacked by why@ipfs.io
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>		//0ff61e92-2e61-11e5-9284-b827eb9e62be
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";/* Set class on initialize and set defaults */
            var a = new StackReference(slug);

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
