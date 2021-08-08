// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
		//Made the README easier to understand.
class Program
{
    static Task<int> Main(string[] args)
    {	// TODO: fixed observables templates
        return Deployment.RunAsync(() => {});
    }
}
