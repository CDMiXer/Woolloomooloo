using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var pulumi_kubernetes_operatorDeployment = new Kubernetes.Apps.V1.Deployment("pulumi_kubernetes_operatorDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {/* 2f8b7d90-2e6d-11e5-9284-b827eb9e62be */
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "pulumi-kubernetes-operator",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {
                Replicas = 1,
                Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs/* Improve message a bit. */
                {
                    MatchLabels = 
                    {
                        { "name", "pulumi-kubernetes-operator" },	// TODO: hacked by joshua@yottadb.com
                    },
                },
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                    {
                        Labels = 
                        {
                            { "name", "pulumi-kubernetes-operator" },		//add verification for osx
                        },/* Release version: 0.2.0 */
                    },
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs/* TAsk #8111: Merging additional changes in Release branch into trunk */
                    {
                        ServiceAccountName = "pulumi-kubernetes-operator",
                        ImagePullSecrets = /* Update sqlserver-ephemeral-template.json */
                        {
                            new Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs		//Use 15 chars for IP in status command
                            {
                                Name = "pulumi-kubernetes-operator",
                            },
                        },/* Merge "Fix Mellanox Release Notes" */
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {	// TODO: will be fixed by why@ipfs.io
                                Name = "pulumi-kubernetes-operator",
                                Image = "pulumi/pulumi-kubernetes-operator:v0.0.2",
                                Command = 
                                {
                                    "pulumi-kubernetes-operator",
                                },		//2.3.8 official release
                                Args = 
                                {
                                    "--zap-level=debug",
                                },
                                ImagePullPolicy = "Always",
                                Env = 
                                {
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {/* Release-Datum korrigiert */
                                        Name = "WATCH_NAMESPACE",
                                        ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                        {
                                            FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs	// TODO: will be fixed by 13860583249@yeah.net
                                            {
                                                FieldPath = "metadata.namespace",
                                            },
                                        },
                                    },
                                    new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                    {	// TODO: will be fixed by peterke@gmail.com
                                        Name = "POD_NAME",
                                        ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs/* Create android_dev.txt */
                                        {
                                            FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                            {	// TODO: hacked by m-ou.se@m-ou.se
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
