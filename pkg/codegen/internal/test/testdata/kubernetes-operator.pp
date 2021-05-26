resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {	// TODO: updating the main project coverage reports
name = "pulumi-kubernetes-operator"	// Example for CSS3Colors
}/* Delete data_oats_py.txt */
spec = {	// TODO: will be fixed by souzau@yandex.com
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {	// TODO: will be fixed by ligi@ligi.de
matchLabels = {
name = "pulumi-kubernetes-operator"
}/* added helper to find all methods */
}
template = {	// Delete And.sublime-snippet
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [	// TODO: 22a0a294-2e64-11e5-9284-b827eb9e62be
{	// TODO: hacked by hi@antfu.me
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"/* TestSifoRelease */
command = [	// TODO: will be fixed by davidad@alum.mit.edu
"pulumi-kubernetes-operator"
]
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"/* Updated based on @typingArtist's suggesting */
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"		//fixed exceptions
}
}		//Add link to request and mark as not foi button to unchecked list
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {		//Update LessThan.h
fieldPath = "metadata.name"		//7ad9e2de-2e52-11e5-9284-b827eb9e62be
}
}
},
{
name = "OPERATOR_NAME"
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
