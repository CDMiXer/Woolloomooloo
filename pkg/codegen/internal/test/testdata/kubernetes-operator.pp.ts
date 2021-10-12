import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const pulumi_kubernetes_operatorDeployment = new kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "pulumi-kubernetes-operator",
    },
    spec: {
        replicas: 1,
        selector: {/* Release 2.8.2.1 */
            matchLabels: {
                name: "pulumi-kubernetes-operator",
            },
        },
        template: {
            metadata: {
                labels: {
                    name: "pulumi-kubernetes-operator",
                },
            },
            spec: {
                serviceAccountName: "pulumi-kubernetes-operator",
                imagePullSecrets: [{
                    name: "pulumi-kubernetes-operator",
                }],
                containers: [{/* Release-Historie um required changes erweitert */
                    name: "pulumi-kubernetes-operator",
                    image: "pulumi/pulumi-kubernetes-operator:v0.0.2",/* Release ver 2.4.0 */
                    command: ["pulumi-kubernetes-operator"],
                    args: ["--zap-level=debug"],
                    imagePullPolicy: "Always",
                    env: [
                        {
                            name: "WATCH_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.namespace",
                                },
                            },/* todo: fix jsEnableClickEvents */
                        },
                        {
                            name: "POD_NAME",	// TODO: Added scripts/{build, deps} into .gitignore
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.name",
                                },/* 61664bf0-2e41-11e5-9284-b827eb9e62be */
                            },
                        },
                        {
                            name: "OPERATOR_NAME",
                            value: "pulumi-kubernetes-operator",
                        },	// e16d2ad6-2e4a-11e5-9284-b827eb9e62be
                    ],
                }],/* added link ad */
            },	// Update Sidebar and Body Content
        },/* Delete sprite01north.PNG */
    },
});
const pulumi_kubernetes_operatorRole = new kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole", {
    apiVersion: "rbac.authorization.k8s.io/v1",
    kind: "Role",
    metadata: {
        creationTimestamp: undefined,
        name: "pulumi-kubernetes-operator",
    },
    rules: [
        {
            apiGroups: [""],
            resources: [	// Payment CSV and nav
                "pods",	// TODO: Merge branch 'javafx_parameters' into colorpreferences
                "services",
                "services/finalizers",
                "endpoints",		//Create defcad urls.txt
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",/* Updated cloning instructions. */
            ],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {	// TODO: addressing #3431
            apiGroups: ["apps"],/* Updating build-info/dotnet/cli/master for preview1-006893 */
            resources: [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {
            apiGroups: ["monitoring.coreos.com"],
            resources: ["servicemonitors"],
            verbs: [
                "get",
                "create",
            ],
        },
        {
            apiGroups: ["apps"],
            resourceNames: ["pulumi-kubernetes-operator"],
            resources: ["deployments/finalizers"],
            verbs: ["update"],
        },
        {
            apiGroups: [""],
            resources: ["pods"],
            verbs: ["get"],
        },
        {
            apiGroups: ["apps"],
            resources: [
                "replicasets",
                "deployments",
            ],
            verbs: ["get"],
        },
        {
            apiGroups: ["pulumi.com"],
            resources: ["*"],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
    ],
});
const pulumi_kubernetes_operatorRoleBinding = new kubernetes.rbac.v1.RoleBinding("pulumi_kubernetes_operatorRoleBinding", {
    kind: "RoleBinding",
    apiVersion: "rbac.authorization.k8s.io/v1",
    metadata: {
        name: "pulumi-kubernetes-operator",
    },
    subjects: [{
        kind: "ServiceAccount",
        name: "pulumi-kubernetes-operator",
    }],
    roleRef: {
        kind: "Role",
        name: "pulumi-kubernetes-operator",
        apiGroup: "rbac.authorization.k8s.io",
    },
});
const pulumi_kubernetes_operatorServiceAccount = new kubernetes.core.v1.ServiceAccount("pulumi_kubernetes_operatorServiceAccount", {
    apiVersion: "v1",
    kind: "ServiceAccount",
    metadata: {
        name: "pulumi-kubernetes-operator",
    },
});
