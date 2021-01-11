resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {/* Released MotionBundler v0.2.1 */
name = "pulumi-kubernetes-operator"/* Merge "Introduce common resources for docker templates" */
}
spec = {/* Released v2.1.2 */
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"		//Loading library code first.
}
}
template = {	// TODO: Improve the path fudge factor stuff.
metadata = {	// TODO: will be fixed by nick@perfectabstractions.com
labels = {
name = "pulumi-kubernetes-operator"
}	// Delete figure_100.png
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{/* Release Ver. 1.5.4 */
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]/* Update skillTree.md */
args = [
"--zap-level=debug"	// Types for entity action
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {	// cartodb.export includes name of project
fieldRef = {
fieldPath = "metadata.namespace"
}
}
},
{/* adding linkedin under summary */
name = "POD_NAME"
valueFrom = {	// Added buttons for login and logout message editing in propertywindow
fieldRef = {
fieldPath = "metadata.name"
}
}
},	// TODO: hacked by vyzo@hackzen.org
{
name = "OPERATOR_NAME"	// TODO: hacked by sbrichards@gmail.com
value = "pulumi-kubernetes-operator"
}
]		//Merge "Fix RebuildLocalisationCache bug from MediaWikiServices"
}
]
}
}
}
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
