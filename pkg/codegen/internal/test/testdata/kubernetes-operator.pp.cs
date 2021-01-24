using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack/* JSDemoApp should be GC in Release too */
{
    public MyStack()/* RuleUtils.mapWithSrc to iterate over the source elements */
    {
        var pulumi_kubernetes_operatorDeployment = new Kubernetes.Apps.V1.Deployment("pulumi_kubernetes_operatorDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs/* Rename future to CNAME */
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {		//Mark the branch with an alpha version number.
                Name = "pulumi-kubernetes-operator",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs/* f38a57ab-2d3e-11e5-8b0b-c82a142b6f9b */
            {	// TODO: Use pytest and flake8 commands.
                Replicas = 1,
                Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
                {	// TODO: More additions to Swedish tsx-file.
                    MatchLabels = 
                    {
                        { "name", "pulumi-kubernetes-operator" },
                    },	// added tests for pooled connections
                },/* Release 0.4.7. */
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {	// TODO: check fcmacro
                    Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                    {	// fix a minor typo in javadoc for variables
                        Labels = 
                        {/* Correção da classe RestResource */
                            { "name", "pulumi-kubernetes-operator" },/* Release 7.7.0 */
                        },
                    },/* updated i/o for genbank files to accept newer versions */
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {
                        ServiceAccountName = "pulumi-kubernetes-operator",
                        ImagePullSecrets = 
                        {		//Update section-overview-home.md
                            new Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs
                            {/* Merge "Tweak webhook middleware for optimization" */
                                Name = "pulumi-kubernetes-operator",
                            },
                        },
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                Name = "pulumi-kubernetes-operator",
                                Image = "pulumi/pulumi-kubernetes-operator:v0.0.2",
                                Command = 
                                {
                                    "pulumi-kubernetes-operator",
                                },
                                Args = 
                                {
                                    "--zap-level=debug",
                                },
                                ImagePullPolicy = "Always",
                                Env = 
                                {
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {
                                        Name = "WATCH_NAMESPACE",
                                        ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                        {
                                            FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                            {
                                                FieldPath = "metadata.namespace",
                                            },
                                        },
                                    },
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {
                                        Name = "POD_NAME",
                                        ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                        {
                                            FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                            {
                                                FieldPath = "metadata.name",
                                            },
                                        },
                                    },
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {
                                        Name = "OPERATOR_NAME",
                                        Value = "pulumi-kubernetes-operator",
                                    },
                                },
                            },
                        },
                    },
                },
            },
        });
        var pulumi_kubernetes_operatorRole = new Kubernetes.Rbac.V1.Role("pulumi_kubernetes_operatorRole", new Kubernetes.Types.Inputs.Rbac.V1.RoleArgs
        {
            ApiVersion = "rbac.authorization.k8s.io/v1",
            Kind = "Role",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                CreationTimestamp = null,
                Name = "pulumi-kubernetes-operator",
            },
            Rules = 
            {
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "",
                    },
                    Resources = 
                    {
                        "pods",
                        "services",
                        "services/finalizers",
                        "endpoints",
                        "persistentvolumeclaims",
                        "events",
                        "configmaps",
                        "secrets",
                    },
                    Verbs = 
                    {
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "apps",
                    },
                    Resources = 
                    {
                        "deployments",
                        "daemonsets",
                        "replicasets",
                        "statefulsets",
                    },
                    Verbs = 
                    {
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "monitoring.coreos.com",
                    },
                    Resources = 
                    {
                        "servicemonitors",
                    },
                    Verbs = 
                    {
                        "get",
                        "create",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "apps",
                    },
                    ResourceNames = 
                    {
                        "pulumi-kubernetes-operator",
                    },
                    Resources = 
                    {
                        "deployments/finalizers",
                    },
                    Verbs = 
                    {
                        "update",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "",
                    },
                    Resources = 
                    {
                        "pods",
                    },
                    Verbs = 
                    {
                        "get",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "apps",
                    },
                    Resources = 
                    {
                        "replicasets",
                        "deployments",
                    },
                    Verbs = 
                    {
                        "get",
                    },
                },
                new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
                {
                    ApiGroups = 
                    {
                        "pulumi.com",
                    },
                    Resources = 
                    {
                        "*",
                    },
                    Verbs = 
                    {
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch",
                    },
                },
            },
        });
        var pulumi_kubernetes_operatorRoleBinding = new Kubernetes.Rbac.V1.RoleBinding("pulumi_kubernetes_operatorRoleBinding", new Kubernetes.Types.Inputs.Rbac.V1.RoleBindingArgs
        {
            Kind = "RoleBinding",
            ApiVersion = "rbac.authorization.k8s.io/v1",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "pulumi-kubernetes-operator",
            },
            Subjects = 
            {
                new Kubernetes.Types.Inputs.Rbac.V1.SubjectArgs
                {
                    Kind = "ServiceAccount",
                    Name = "pulumi-kubernetes-operator",
                },
            },
            RoleRef = new Kubernetes.Types.Inputs.Rbac.V1.RoleRefArgs
            {
                Kind = "Role",
                Name = "pulumi-kubernetes-operator",
                ApiGroup = "rbac.authorization.k8s.io",
            },
        });
        var pulumi_kubernetes_operatorServiceAccount = new Kubernetes.Core.V1.ServiceAccount("pulumi_kubernetes_operatorServiceAccount", new Kubernetes.Types.Inputs.Core.V1.ServiceAccountArgs
        {
            ApiVersion = "v1",
            Kind = "ServiceAccount",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "pulumi-kubernetes-operator",
            },
        });
    }

}
