// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: optimization for "toString()"
using System.Threading.Tasks;/* Merge "Fix order of arguments in assertEqual - Part3" */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => {});
    }
}
