resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33	// a00ab7a4-2e6b-11e5-9284-b827eb9e62be
replicas = 1
selector = {/* Fix @Override in Eclipse. */
matchLabels = {
name = "pulumi-kubernetes-operator"/* Release-Notes f. Bugfix-Release erstellt */
}
}
template = {
metadata = {	// Create README01.md
labels = {
name = "pulumi-kubernetes-operator"		//open a dialog on login error #28
}
}
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
command = [		//Add greenkeeper-lockfile to CI
"pulumi-kubernetes-operator"
]
args = [
"--zap-level=debug"
]/* Add Reserve / Decode */
imagePullPolicy = "Always"		//- fixed a potential problem with Playlist
env = [/* Issue #44 Release version and new version as build parameters */
{
name = "WATCH_NAMESPACE"
valueFrom = {	// Merge "Add not set value to ports filtering in selector"
fieldRef = {
fieldPath = "metadata.namespace"/* add code to reselect an app in the list view after a model refresh */
}
}
},
{	// Removed libSBOLj from local maven repo.
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
}	// TODO: will be fixed by 13860583249@yeah.net
}
},
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}		//Update message.
]
}/* Removing some duplicated code in IncludeFlattener */
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
