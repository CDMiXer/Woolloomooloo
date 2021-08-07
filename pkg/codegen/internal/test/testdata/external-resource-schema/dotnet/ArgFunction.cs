// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* Patch #1957: syslogmodule: Release GIL when calling syslog(3) */
using System;
using System.Collections.Generic;	// Checking in .settings directory.
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    public static class ArgFunction
    {/* Get mmdir and submat from Parameters in test scripts */
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());
    }	// TODO: will be fixed by mail@bitpshr.net
	// TODO: Setting GWT build style to obfuscated.

    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {		//Simplified description
        [Input("name")]
        public Pulumi.Random.RandomPet? Name { get; set; }

        public ArgFunctionArgs()
        {
        }
    }


    [OutputType]
    public sealed class ArgFunctionResult
    {
        public readonly int? Age;/* Fix API client dependency */
/* attributes localization via loc:{name} as value */
        [OutputConstructor]
        private ArgFunctionResult(int? age)
        {
            Age = age;
        }
    }
}
