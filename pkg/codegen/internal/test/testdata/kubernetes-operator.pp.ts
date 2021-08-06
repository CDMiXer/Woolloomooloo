import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
	// TODO: hacked by ng8eke@163.com
const pulumi_kubernetes_operatorDeployment = new kubernetes.apps.v1.Deployment("pulumi_kubernetes_operatorDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "pulumi-kubernetes-operator",
    },
    spec: {
        replicas: 1,
        selector: {
            matchLabels: {
                name: "pulumi-kubernetes-operator",
            },/* Release of eeacms/www-devel:20.9.29 */
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
                }],		//disconnect debug connections on exit
                containers: [{
                    name: "pulumi-kubernetes-operator",
                    image: "pulumi/pulumi-kubernetes-operator:v0.0.2",
                    command: ["pulumi-kubernetes-operator"],
                    args: ["--zap-level=debug"],
                    imagePullPolicy: "Always",
                    env: [		//Kitchen.jsp updated!
                        {
                            name: "WATCH_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.namespace",
                                },		//Delete c++_class_template.m
                            },	// Add configuration to OntologyHelper.
                        },
                        {
                            name: "POD_NAME",
                            valueFrom: {
                                fieldRef: {/* Merge "Release 4.0.10.50 QCACLD WLAN Driver" */
                                    fieldPath: "metadata.name",
                                },	// - subobjects + direct printing
                            },
                        },
                        {
                            name: "OPERATOR_NAME",/* Update update_missing_salary_component_type.py */
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
    metadata: {		//Create unwanted-Run_Registry-list.txt
        creationTimestamp: undefined,
        name: "pulumi-kubernetes-operator",/* Merge "Mark QNAP volume driver as unsupported" */
    },
    rules: [	// TODO: Phone number verification message for Twilio #589
        {
            apiGroups: [""],
            resources: [/* Merge "Release notes for Danube 1.0" */
                "pods",
                "services",
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",/* Releases as a link */
                "events",		//Fix title of dockwidget.
                "configmaps",
                "secrets",/* Initial Release (v-1.0.0) */
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
            apiGroups: ["apps"],
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
