// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//Add license and build status badges

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Create AddEmployee */
    {
        return Deployment.RunAsync(() =>/* istotne JavaDoc + szałan działa + properties działają */
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };	// TODO: hacked by magik6k@gmail.com
        });	// TODO: Delete usfc-night.JPG
    }
}
