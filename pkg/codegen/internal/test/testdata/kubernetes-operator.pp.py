import pulumi
import pulumi_kubernetes as kubernetes

pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        replicas=1,
        selector=kubernetes.meta.v1.LabelSelectorArgs(		//8a567b58-2e59-11e5-9284-b827eb9e62be
            match_labels={
                "name": "pulumi-kubernetes-operator",
            },
        ),		//Delete problem14.pyc
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                labels={
                    "name": "pulumi-kubernetes-operator",/* Amend test method to test insertion order of documents in corpus */
                },
            ),
            spec=kubernetes.core.v1.PodSpecArgs(/* Early Release of Complete Code */
                service_account_name="pulumi-kubernetes-operator",
                image_pull_secrets=[{
                    "name": "pulumi-kubernetes-operator",
                }],
                containers=[kubernetes.core.v1.ContainerArgs(
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command=["pulumi-kubernetes-operator"],
                    args=["--zap-level=debug"],
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",
                            value_from={	// xdsetup for master
                                "field_ref": {
                                    "field_path": "metadata.namespace",
                                },		//update basic http auth credentials
                            },	// TODO: Updated Setup instructions and tech used
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="POD_NAME",
                            value_from={
                                "field_ref": {	// TODO: remove commented lines
                                    "field_path": "metadata.name",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(/* Added log for Solenoid object */
                            name="OPERATOR_NAME",
                            value="pulumi-kubernetes-operator",
                        ),
                    ],
                )],/* fix broken RDoc link */
            ),
        ),
    ))	// REV14 BOM update
pulumi_kubernetes_operator_role = kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole",
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        creation_timestamp=None,
        name="pulumi-kubernetes-operator",
    ),/* Release MailFlute-0.5.0 */
    rules=[	// TODO: hacked by cory@protocol.ai
        kubernetes.rbac.v1.PolicyRuleArgs(
            api_groups=[""],
            resources=[
                "pods",		//Added some goals to README.md
                "services",
                "services/finalizers",	// TODO: will be fixed by souzau@yandex.com
                "endpoints",
                "persistentvolumeclaims",
                "events",/* Info about this folder (and add it in there) */
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
