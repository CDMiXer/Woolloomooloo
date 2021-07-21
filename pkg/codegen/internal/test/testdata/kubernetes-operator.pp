resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {/* Put Initial Release Schedule */
apiVersion = "apps/v1"/* Release 1.0.1. */
kind = "Deployment"
metadata = {		//Add more tests and business code for time-tracker
name = "pulumi-kubernetes-operator"		//Release 0.8 Alpha
}	// TODO: begin implementation of the control selection strategy
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {	// TODO: round the duration, probe
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"/* Sync ChangeLog and ReleaseNotes */
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [		//Allow a custom box to be specified for the colorbar
"pulumi-kubernetes-operator"
]/* Task #7064: Imported Release 2.8 fixes (AARTFAAC and DE609 changes) */
args = [		//updating combi stuff
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {
fieldRef = {/* Added missng include directory to Xcode project for Release build. */
fieldPath = "metadata.namespace"/* Release version 1.1.3 */
}
}
},
{	// Tela de vendas atualizada com bd e xml
name = "POD_NAME"
valueFrom = {
fieldRef = {/* #714: MapTileRastered can set custom raster line number. */
fieldPath = "metadata.name"
}
}	// TODO: Fix typos in irc_sprintf documentation
},
{
name = "OPERATOR_NAME"/* Merge "Release 3.2.3.419 Prima WLAN Driver" */
value = "pulumi-kubernetes-operator"
}
]
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
