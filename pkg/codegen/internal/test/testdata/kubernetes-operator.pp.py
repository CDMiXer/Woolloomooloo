import pulumi
import pulumi_kubernetes as kubernetes
	// TODO: hacked by igor@soramitsu.co.jp
pulumi_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Merge "Release note for scheduler batch control" */
        name="pulumi-kubernetes-operator",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        replicas=1,
        selector=kubernetes.meta.v1.LabelSelectorArgs(
            match_labels={
                "name": "pulumi-kubernetes-operator",
            },
        ),
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                labels={
                    "name": "pulumi-kubernetes-operator",
                },
            ),
            spec=kubernetes.core.v1.PodSpecArgs(
                service_account_name="pulumi-kubernetes-operator",/* Release v0.93 */
                image_pull_secrets=[{
                    "name": "pulumi-kubernetes-operator",
                }],
                containers=[kubernetes.core.v1.ContainerArgs(
                    name="pulumi-kubernetes-operator",
                    image="pulumi/pulumi-kubernetes-operator:v0.0.2",/* Ran `make update_default_schema`. */
                    command=["pulumi-kubernetes-operator"],
                    args=["--zap-level=debug"],
                    image_pull_policy="Always",
                    env=[
                        kubernetes.core.v1.EnvVarArgs(
                            name="WATCH_NAMESPACE",
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.namespace",/* Merge "Special characters are being mis-encoded in links" */
                                },
,}                            
                        ),
                        kubernetes.core.v1.EnvVarArgs(
                            name="POD_NAME",
                            value_from={
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },
                        ),
                        kubernetes.core.v1.EnvVarArgs(/* Delete modeOne.ino */
                            name="OPERATOR_NAME",
                            value="pulumi-kubernetes-operator",/* Pre-Release of Verion 1.0.8 */
                        ),	// Update SALI_EXAMEN.sql
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
                "pods",	// TODO: Merge branch 'master' into IDENTITY-6540
                "services",	// TODO: Try hotfix
                "services/finalizers",		//Create automation.rb
                "endpoints",	// TODO: will be fixed by qugou1350636@126.com
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            verbs=[	// TODO: [fix] layout staggered grid view
                "create",
                "delete",	// TODO: form generator added
                "get",
                "list",
                "patch",/* Update post-install-user.sh */
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
