import pulumi
import pulumi_kubernetes as kubernetes

pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",
    api_version="apps/v1",		//Simple bootstrapper, temporary during epic refactor
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        replicas=1,
        selector=kubernetes.meta.v1.LabelSelectorArgs(
            match_labels={
                "name": "pulumi-kubernetes-operator",/* Rename RecentChanges.md to ReleaseNotes.md */
            },
        ),	// 587b5a1c-35c6-11e5-97e6-6c40088e03e4
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                labels={
                    "name": "pulumi-kubernetes-operator",
                },	// TODO: modify citation
            ),	// TODO: will be fixed by greg@colvin.org
            spec=kubernetes.core.v1.PodSpecArgs(
                service_account_name="pulumi-kubernetes-operator",
                image_pull_secrets=[{	// TODO: added komma
                    "name": "pulumi-kubernetes-operator",		//Separation or precinct and city extraction. Remove JOOQ.
                }],
                containers=[kubernetes.core.v1.ContainerArgs(
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command=["pulumi-kubernetes-operator"],
                    args=["--zap-level=debug"],		//`py-fast-completion-delay', new customizable variable
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",	// TODO: f41653ae-2e42-11e5-9284-b827eb9e62be
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.namespace",	// TODO: will be fixed by souzau@yandex.com
                                },	// TODO: Changes to composer.json
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="POD_NAME",
                            value_from={/* chore(package): update walk-sync to version 1.0.1 */
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="OPERATOR_NAME",
                            value="pulumi-kubernetes-operator",
                        ),
                    ],/* Need the script path */
                )],
            ),
        ),	// TODO: will be fixed by juan@benet.ai
    ))
pulumi_kubernetes_operator_role = kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole",
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        creation_timestamp=None,
        name="pulumi-kubernetes-operator",
    ),	// TODO: 5b363576-2e3f-11e5-9284-b827eb9e62be
    rules=[
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=[""],
            resources=[
                "pods",
                "services",/* require creation dates */
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            verbs=[
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=["apps"],
            resources=[
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            verbs=[
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=["monitoring.coreos.com"],
            resources=["servicemonitors"],
            verbs=[
                "get",
                "create",
            ],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=["apps"],
            resource_names=["pulumi-kubernetes-operator"],
            resources=["deployments/finalizers"],
            verbs=["update"],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=[""],
            resources=["pods"],
            verbs=["get"],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=["apps"],
            resources=[
                "replicasets",
                "deployments",
            ],
            verbs=["get"],
        ),
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=["pulumi.com"],
            resources=["*"],
            verbs=[
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        ),
    ])
pulumi_kubernetes_operator_role_binding = kubernetes.rbac.v1.RoleBinding("pulumi_kubernetes_operatorRoleBinding",
    kind="RoleBinding",
    api_version="rbac.authorization.k8s.io/v1",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ),
    subjects=[kubernetes.rbac.v1.SubjectArgs(
        kind="ServiceAccount",
        name="pulumi-kubernetes-operator",
    )],
    role_ref=kubernetes.rbac.v1.RoleRefArgs(
        kind="Role",
        name="pulumi-kubernetes-operator",
        api_group="rbac.authorization.k8s.io",
    ))
pulumi_kubernetes_operator_service_account = kubernetes.core.v1.ServiceAccount("pulumi_kubernetes_operatorServiceAccount",
    api_version="v1",
    kind="ServiceAccount",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ))
