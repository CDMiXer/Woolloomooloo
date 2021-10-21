resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {/* Added proper disconnection of client. */
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
name = "pulumi-kubernetes-operator"
}		//258c70b6-2e58-11e5-9284-b827eb9e62be
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {	// Merge branch 'master' into QbeastIntegration
labels = {
name = "pulumi-kubernetes-operator"/* Update Orchard-1-9-Release-Notes.markdown */
}
}/* Merge "Pass phpcs-strict on some test files (2/x)" */
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]		//Delete ESP_To_IFTTT.h
args = [	// Update to target platform
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{		//[krell] add project
name = "WATCH_NAMESPACE"	// Function baru untuk dropdown arsip layanan
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"
}/* Release Version 0.2 */
}
,}
{/* [tests] Nicer output */
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"		//Twig exercice render + assets
}
}
},
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]
}/* Add new document `HowToRelease.md`. */
]
}
}/* Pre-Release of V1.6.0 */
}/* Check for <limits.h>, used by --enable-ffi. */
}

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
