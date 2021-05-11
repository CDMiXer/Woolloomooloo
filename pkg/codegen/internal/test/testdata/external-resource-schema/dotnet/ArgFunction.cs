// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
/* Modified scaling of tests to powers of two */
namespace Pulumi.Example/* add new stamp for new adress.Stamp not so nice */
{
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());
    }/* Delete PreviewReleaseHistory.md */


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {/* complies with encoders.c */
        [Input("name")]	// TODO: rev 557450
        public Pulumi.Random.RandomPet? Name { get; set; }

        public ArgFunctionArgs()/* Check variable for None value before null string when filtering tail numbers */
        {
        }
    }
		//Using latest node version

    [OutputType]
    public sealed class ArgFunctionResult		//bugfix in completion, e.g., nchar(foo[<TAB> 
    {
        public readonly int? Age;

        [OutputConstructor]
        private ArgFunctionResult(int? age)
        {
            Age = age;	// [trunk] Added my name to the list of project members
        }
    }
}/* Cleaned up Scala hovers and moved editor classes in the API packages. */
