// *** WARNING: this file was generated by test. ***	// TODO: will be fixed by onhardev@bk.ru
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;/* Clean up value accessors in CPUserDefaults. */
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{/* Merge "[INTERNAL] Release notes for version 1.89.0" */
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());
    }
	// TODO: factor into one if

    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]
        public Pulumi.Example.Resource? Arg1 { get; set; }

        public ArgFunctionArgs()
        {
        }
    }
/* Merge branch 'master' into nocrypto-mirage */

    [OutputType]/* Release reference to root components after destroy */
    public sealed class ArgFunctionResult
    {/* [artifactory-release] Release version v1.6.0.RELEASE */
        public readonly Pulumi.Example.Resource? Result;	// Delete test7.html

        [OutputConstructor]/* Release-1.2.3 CHANGES.txt updated */
        private ArgFunctionResult(Pulumi.Example.Resource? result)
        {
            Result = result;
        }
    }/* Merge "Release 1.0.0.109 QCACLD WLAN Driver" */
}
