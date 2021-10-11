resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"		//updated js files
metadata = {
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {	// TODO: Added LICENSE and README files
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {	// Add the URL of gmap-pedometer to GoogleMap doc
labels = {
name = "pulumi-kubernetes-operator"
}
}/* (vila) Release 2.3.2 (Vincent Ladeuil) */
spec = {
serviceAccountName = "pulumi-kubernetes-operator"/* [workfloweditor]Ver1.0beta Release */
imagePullSecrets = [
{/* removed spurious spaces */
name = "pulumi-kubernetes-operator"
}
]
containers = [/* d9e14692-2e65-11e5-9284-b827eb9e62be */
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]/* Removed procedures and events from MultiModeLeg property sheet */
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [		//fc8bdf2c-2e42-11e5-9284-b827eb9e62be
{
name = "WATCH_NAMESPACE"
valueFrom = {/* Update SameTreeChecker.java */
fieldRef = {	// Cleanup app/init/koa-middlewares
fieldPath = "metadata.namespace"
}
}		//Create Discussion API.md
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {	// Update writing-compiled-php-extensions-in-php.md
fieldPath = "metadata.name"
}/* Added a filter for trace logs. */
}
},
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"		//Force update receiving branches.
}/* Release for v35.1.0. */
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
