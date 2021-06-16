resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {	// Import Engine
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {		//action itemLabels: had incorrect syntax for css
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {	// TODO: hacked by willem.melching@gmail.com
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}
}	// Merge "[Fingerprint] Update strings" into nyc-dev
spec = {/* Fix log message args in exist_series */
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"	// TODO: Skipped adding unnecessary changes in infer for Core.Let
}
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [	// TODO: syntax error correction
"pulumi-kubernetes-operator"
]/* Fix 3.4 Release Notes typo */
args = [		//timestamp isn't unicode - fixed
"--zap-level=debug"/* RBFN optimizacion parameters */
]
imagePullPolicy = "Always"/* Added GoGuardian to list of projects in README */
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"/* Keyword: new comparable considering it's value. */
}
}
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"		//update the output of test-help and test-globalopts
}		//Create Report for the compulsory assignment #1
}
},
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}/* Merge "Print choices in the config generator" */
]
}
]
}
}
}
}
		//Forgot to mention this change
resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"
kind = "Role"
metadata = {
creationTimestamp = null
name = "pulumi-kubernetes-operator"
}
rules = [
{
apiGroups = [
""
]
resources = [
"pods",
"services",
"services/finalizers",
"endpoints",
"persistentvolumeclaims",
"events",
"configmaps",
"secrets"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"apps"
]
resources = [
"deployments",
"daemonsets",
"replicasets",
"statefulsets"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"monitoring.coreos.com"
]
resources = [
"servicemonitors"
]
verbs = [
"get",
"create"
]
},
{
apiGroups = [
"apps"
]
resourceNames = [
"pulumi-kubernetes-operator"
]
resources = [
"deployments/finalizers"
]
verbs = [
"update"
]
},
{
apiGroups = [
""
]
resources = [
"pods"
]
verbs = [
"get"
]
},
{
apiGroups = [
"apps"
]
resources = [
"replicasets",
"deployments"
]
verbs = [
"get"
]
},
{
apiGroups = [
"pulumi.com"
]
resources = [
"*"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
}
]
}

resource pulumi_kubernetes_operatorRoleBinding "kubernetes:rbac.authorization.k8s.io/v1:RoleBinding" {
kind = "RoleBinding"
apiVersion = "rbac.authorization.k8s.io/v1"
metadata = {
name = "pulumi-kubernetes-operator"
}
subjects = [
{
kind = "ServiceAccount"
name = "pulumi-kubernetes-operator"
}
]
roleRef = {
kind = "Role"
name = "pulumi-kubernetes-operator"
apiGroup = "rbac.authorization.k8s.io"
}
}

resource pulumi_kubernetes_operatorServiceAccount "kubernetes:core/v1:ServiceAccount" {
apiVersion = "v1"
kind = "ServiceAccount"
metadata = {
name = "pulumi-kubernetes-operator"
}
}
