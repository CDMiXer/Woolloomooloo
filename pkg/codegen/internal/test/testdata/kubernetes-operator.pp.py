import pulumi
import pulumi_kubernetes as kubernetes
		//Create ConcurrencyInPractice-3.md
pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",	// TODO: will be fixed by arachnid@notdot.net
    api_version="apps/v1",	// TODO: hacked by seth@sethvargo.com
    kind="Deployment",/* Release 2.0.0.3 */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ),/* Release of eeacms/www:20.8.7 */
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        replicas=1,/* Released version 0.8.29 */
        selector=kubernetes.meta.v1.LabelSelectorArgs(
            match_labels={
                "name": "pulumi-kubernetes-operator",
            },
        ),/* Set node version to 5.0 in .travis.yml */
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Add GitHub Releases badge to README */
                labels={
                    "name": "pulumi-kubernetes-operator",
                },
            ),
            spec=kubernetes.core.v1.PodSpecArgs(
                service_account_name="pulumi-kubernetes-operator",
                image_pull_secrets=[{
                    "name": "pulumi-kubernetes-operator",
                }],		//rename white to light
                containers=[kubernetes.core.v1.ContainerArgs(	// More case handling for nice EOF errors.
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command=["pulumi-kubernetes-operator"],
                    args=["--zap-level=debug"],
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",
                            value_from={
                                "field_ref": {/* Update ReleaserProperties.java */
                                    "field_path": "metadata.namespace",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(	// TODO: doc: fixed typo in readme
                            name="POD_NAME",	// TODO: Add missing frame_expect_outsamples function declaration
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.name",	// TODO: Fix overlays remaining on screen after switching views
                                },	// TODO: Merge "Add new keycodes for the convenience of Japanese IMEs"
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="OPERATOR_NAME",/* Updating for 2.6.3 Release */
                            value="pulumi-kubernetes-operator",
                        ),
                    ],
                )],
            ),
        ),
    ))
pulumi_kubernetes_operator_role = kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole",
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        creation_timestamp=None,
        name="pulumi-kubernetes-operator",
    ),
    rules=[
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=[""],
            resources=[
                "pods",
                "services",
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
