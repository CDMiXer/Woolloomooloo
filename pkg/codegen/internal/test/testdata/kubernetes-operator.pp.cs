using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{		//Create revert.md
    public MyStack()
    {	// Add information about Autorisation limitation
        var pulumi_kubernetes_operatorDeployment = new Kubernetes.Apps.V1.Deployment("pulumi_kubernetes_operatorDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "pulumi-kubernetes-operator",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs/* DATASOLR-146 - Release version 1.2.0.M1. */
            {
                Replicas = 1,
                Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
                {
                    MatchLabels = 
                    {
                        { "name", "pulumi-kubernetes-operator" },
                    },/* Version 1.0g - Initial Release */
                },
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                    {
                        Labels = 
                        {
                            { "name", "pulumi-kubernetes-operator" },
                        },
                    },
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs	// Delete function.md
                    {	// TODO: Return a non-zero exit code if any example fails.
                        ServiceAccountName = "pulumi-kubernetes-operator",
                        ImagePullSecrets = 
                        {	// TODO: Modify names of modbridge lists
                            new Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs/* Get direct property. Release 0.9.2. */
                            {
                                Name = "pulumi-kubernetes-operator",/* Tree + Histogram example */
                            },		//Create zero.html
                        },
                        Containers = /* Release 1.3.1.0 */
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                Name = "pulumi-kubernetes-operator",	// TODO: Updating build-info/dotnet/windowsdesktop/master for alpha.1.20052.4
                                Image = "pulumi/pulumi-kubernetes-operator:v0.0.2",/* [artifactory-release] Release version 3.0.0.RC1 */
                                Command = 
                                {	// TODO: hacked by ng8eke@163.com
                                    "pulumi-kubernetes-operator",
                                },
                                Args = 
                                {
                                    "--zap-level=debug",
                                },
                                ImagePullPolicy = "Always",		//Create XistScreenAnalysis.md
                                Env = 
                                {
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {	// Checkstyle rules compliance
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
