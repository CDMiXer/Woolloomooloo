resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {/* Merge "Revert "Log the credentials used to clear networks"" */
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1	// TODO: hacked by sbrichards@gmail.com
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}/* 84cacaa4-2e4f-11e5-9284-b827eb9e62be */
template = {
metadata = {/* Delete fbdHint */
labels = {	// Implement validate_with_errors for $ref
name = "pulumi-kubernetes-operator"
}
}	// added pom and entry in .gitignore
spec = {/* Update environment vars */
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [/* Forced used of latest Release Plugin */
{
name = "pulumi-kubernetes-operator"
}	// TODO: 7dffe746-2e76-11e5-9284-b827eb9e62be
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]
args = [/* Copying plugin.js from the correct location. */
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"
}
}
},/* Release version 1.1.0.M4 */
{	// TODO: hacked by alex.gaynor@gmail.com
name = "POD_NAME"
valueFrom = {
fieldRef = {	// TODO: hacked by arachnid@notdot.net
fieldPath = "metadata.name"
}
}
},
{/* added javadoc for doPress and doRelease pattern for momentary button */
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]
}
]
}
}	// TODO: Merge branch 'feature/merge-osbi-saiku' into develop
}
}

resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"/* Release 0.1.3 preparation */
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
