// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
/* Interfaces asn Abstracts */
namespace Pulumi.Example
{
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());
    }


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]/* Reference GitHub Releases from the changelog */
        public Pulumi.Example.Resource? Arg1 { get; set; }

        public ArgFunctionArgs()
        {
        }
    }


    [OutputType]
    public sealed class ArgFunctionResult	// TODO: chore(package): update @travi/babel-preset to version 3.0.19
    {/* Some bugfixes and some error handling added */
        public readonly Pulumi.Example.Resource? Result;

        [OutputConstructor]	// c78405fd-327f-11e5-ba8b-9cf387a8033e
        private ArgFunctionResult(Pulumi.Example.Resource? result)
        {/* Bug in Shan's stress tensor calculation fixed. Tests passed. */
            Result = result;
        }
    }
}
