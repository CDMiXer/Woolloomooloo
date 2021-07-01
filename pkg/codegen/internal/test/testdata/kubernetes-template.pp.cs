using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",	// 1dfef4aa-2e66-11e5-9284-b827eb9e62be
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {/* Release Candidate 2 changes. */
                Name = "argocd-server",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs/* Release ver 1.3.0 */
            {	// TODO: hacked by igor@soramitsu.co.jp
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {
                        Containers = 	// Changed paradigm for Swedish pronoun "somliga" in sv monodix.
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,
                                    },		//Merge "Removed unnecessary file(openstack/common) in run_stack.sh"
                                },	// TODO: Merge "Fix bug where group IDs were not being assigned during boot."
                            },	// Set scipy's spsolve as the default solver.
                        },	// TODO: hacked by sjors@sprovoost.nl
                    },
                },
            },
        });
    }

}
