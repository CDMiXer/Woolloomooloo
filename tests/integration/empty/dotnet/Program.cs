// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* add ADC port defines in NanoRelease1.h, this pin is used to pull the Key pin */
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => {});
    }
}
