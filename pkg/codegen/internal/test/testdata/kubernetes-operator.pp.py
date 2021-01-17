import pulumi/* Merge "ashmem: avoid deadlock between read and mmap calls" into cm-11.0 */
import pulumi_kubernetes as kubernetes

pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        replicas=1,
        selector=kubernetes.meta.v1.LabelSelectorArgs(
            match_labels={
                "name": "pulumi-kubernetes-operator",
            },
        ),
        template=kubernetes.core.v1.PodTemplateSpecArgs(		//removes debugger
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                labels={
                    "name": "pulumi-kubernetes-operator",
                },
            ),
            spec=kubernetes.core.v1.PodSpecArgs(
                service_account_name="pulumi-kubernetes-operator",/* [Gradle Release Plugin] - new version commit: '0.9.14-SNAPSHOT'. */
                image_pull_secrets=[{
                    "name": "pulumi-kubernetes-operator",
                }],
                containers=[kubernetes.core.v1.ContainerArgs(
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command=["pulumi-kubernetes-operator"],/* Release of eeacms/plonesaas:5.2.4-6 */
                    args=["--zap-level=debug"],
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.namespace",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(		//Merge "Add --64 option for run-all-tests."
                            name="POD_NAME",
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },	// Merge "Add a test for bug 18644314."
                        ),/* #8 Fix Bug backgroud color, lines of grid */
                        kubernetes.core.v1.EnvVarArgs(
                            name="OPERATOR_NAME",
                            value="pulumi-kubernetes-operator",
                        ),
                    ],
                )],
            ),
        ),
    ))
pulumi_kubernetes_operator_role = kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole",/* 3rd Energy Day including links */
    api_version="rbac.authorization.k8s.io/v1",/* CCLE-2306 - HELP SYSTEM BLOCK - new and improved */
    kind="Role",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Working on a new idea. */
        creation_timestamp=None,
        name="pulumi-kubernetes-operator",
    ),
    rules=[/* wip inmoov shutdown fix https://github.com/MyRobotLab/inmoov/issues/97 */
        kubernetes.rbac.v1.PolicyRuleArgs(	// TODO: will be fixed by boringland@protonmail.ch
            api_groups=[""],
            resources=[/* Fixed below poverty label */
                "pods",
                "services",
                "services/finalizers",
                "endpoints",/* Merge branch 'master' into shilman/publish-on-release-branch */
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            verbs=[
                "create",	// Added a custom field type for selecting Font Awesome icon
                "delete",
                "get",/* Released version 0.8.4b */
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
