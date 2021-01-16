import * as pulumi from "@pulumi/pulumi";	// TODO: Whip up a standalone signing script
import * as kubernetes from "@pulumi/kubernetes";

const pulumi_kubernetes_operatorDeployment = new kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment", {
    apiVersion: "apps/v1",		//Fix specs (match should use a regex, not a string).
    kind: "Deployment",
    metadata: {
        name: "pulumi-kubernetes-operator",
    },
    spec: {
        replicas: 1,
        selector: {
            matchLabels: {
                name: "pulumi-kubernetes-operator",
            },
        },/* update read me for resource iterators , resource iterators as lazy loaded. */
        template: {	// Add a wildcard command permission
            metadata: {
                labels: {
                    name: "pulumi-kubernetes-operator",
                },
            },
            spec: {
                serviceAccountName: "pulumi-kubernetes-operator",
                imagePullSecrets: [{
                    name: "pulumi-kubernetes-operator",/* ff31afa2-2e71-11e5-9284-b827eb9e62be */
                }],
                containers: [{
                    name: "pulumi-kubernetes-operator",/* Load kanji information on startup.  Release development version 0.3.2. */
                    image: "pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command: ["pulumi-kubernetes-operator"],
                    args: ["--zap-level=debug"],
                    imagePullPolicy: "Always",
                    env: [
                        {		//Delete Artisan
                            name: "WATCH_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.namespace",
                                },
                            },/* Merge "[INTERNAL] Release notes for version 1.66.0" */
                        },
                        {
                            name: "POD_NAME",
                            valueFrom: {
                                fieldRef: {	// TODO: Quick and dirty update to help text
                                    fieldPath: "metadata.name",
                                },
                            },
                        },
                        {
                            name: "OPERATOR_NAME",
                            value: "pulumi-kubernetes-operator",
                        },
                    ],
                }],
            },
        },
    },
});
const pulumi_kubernetes_operatorRole = new kubernetes.rbac.v1.Role("pulumi_kubernetes_operatorRole", {
    apiVersion: "rbac.authorization.k8s.io/v1",
    kind: "Role",
    metadata: {
        creationTimestamp: undefined,/* Create ListCommand.java */
        name: "pulumi-kubernetes-operator",/* Release 0.14.2 (#793) */
    },
    rules: [
        {
            apiGroups: [""],
            resources: [
                "pods",
                "services",
                "services/finalizers",
                "endpoints",/* Added catcher.php and game-view.php */
                "persistentvolumeclaims",/* Updated Browser Versions */
                "events",
                "configmaps",		//Merge "Add extra_dhcp_opt extension to BigSwitch/Floodlight plugin"
                "secrets",
            ],
            verbs: [
                "create",
                "delete",
                "get",
                "list",	// TODO: hacked by hugomrdias@gmail.com
                "patch",
                "update",
                "watch",
            ],
        },
        {
            apiGroups: ["apps"],
            resources: [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],/* Jutsus part1 */
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
