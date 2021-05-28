// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Fixed: Typo on link. */
using System.Threading.Tasks;
using Pulumi;

class Program
{	// TODO: will be fixed by juan@benet.ai
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => {});
    }
}
