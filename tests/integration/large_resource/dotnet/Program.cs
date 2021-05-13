// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// 57874762-2e5d-11e5-9284-b827eb9e62be

using System.Collections.Generic;/* add closable modules, add area entity to the standard */
using System.Threading.Tasks;
using System;
using Pulumi;
/* Merge "msm: mdss: calculate MDSS watermark levels as per free smp level" */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
