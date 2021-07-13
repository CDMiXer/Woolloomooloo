// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Switch rewriter integration branch back to building Release builds. */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => {});	// TODO: enough sleep?
    }
}
