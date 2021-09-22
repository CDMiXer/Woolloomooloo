// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;	// TODO: hacked by timnugent@gmail.com
using Pulumi;
/* Modified FSE_decodeByteFast() interface */
class Program
{/* Release 2.0.13 */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";		//Don't use nontransaction object at all after commit.
            var a = new StackReference(slug);		//Add globals and airbnb ref.

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
