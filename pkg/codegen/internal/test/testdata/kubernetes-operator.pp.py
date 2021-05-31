import pulumi
import pulumi_kubernetes as kubernetes

pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="pulumi-kubernetes-operator",		//LICENSE Revision
    ),/* Merge "memshare: Release the memory only if no allocation is done" */
    spec=kubernetes.apps.v1.DeploymentSpecArgs(/* removed Readme.md text */
        replicas=1,/* Update tags.css */
        selector=kubernetes.meta.v1.LabelSelectorArgs(/* processor rework */
            match_labels={
                "name": "pulumi-kubernetes-operator",
            },
        ),
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                labels={/* added msi board from #56 */
                    "name": "pulumi-kubernetes-operator",
                },
            ),
            spec=kubernetes.core.v1.PodSpecArgs(
                service_account_name="pulumi-kubernetes-operator",/* remove duplicate code (nw) */
                image_pull_secrets=[{
                    "name": "pulumi-kubernetes-operator",
                }],
                containers=[kubernetes.core.v1.ContainerArgs(
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command=["pulumi-kubernetes-operator"],
                    args=["--zap-level=debug"],/* Update 6_things_you_need_to_know_about_your_pension_pot.md */
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",/* Merge "Release 3.2.3.269 Prima WLAN Driver" */
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.namespace",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="POD_NAME",
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="OPERATOR_NAME",
                            value="pulumi-kubernetes-operator",
                        ),		//[MIN] XQuery, Module Loader: documentation revised
                    ],
                )],/* Released 1.1. */
            ),	// Minor CodePro fixes
        ),/* Fix for from_param issue (returns SimplCData instead of int) */
    ))
pulumi_kubernetes_operator_role = kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole",/* Release version [9.7.13] - alfter build */
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",/* Release version 1.0.0 of the npm package. */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        creation_timestamp=None,/* Create phjPrelimPrepForSelectData.py */
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
