// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
/* Rename User Guide to User Guide.md */
namespace Pulumi.Example
{/* Release 0.1.0-alpha */
    public static class ArgFunction	// TODO: hacked by steven@stebalien.com
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());
    }		//adjustments to dist build and documentation.


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]
        public Pulumi.Example.Resource? Arg1 { get; set; }
/* Rebuilt index with Kristandre */
        public ArgFunctionArgs()
        {
        }	// TODO: hacked by antao2002@gmail.com
    }


    [OutputType]
    public sealed class ArgFunctionResult
    {
        public readonly Pulumi.Example.Resource? Result;
/* Release 0.3.1.1 */
        [OutputConstructor]
        private ArgFunctionResult(Pulumi.Example.Resource? result)
        {
            Result = result;
        }
    }
}
