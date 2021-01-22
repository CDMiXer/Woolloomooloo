resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"	// TODO: will be fixed by alex.gaynor@gmail.com
metadata = {
name = "pulumi-kubernetes-operator"/* Update Tests.cpp */
}
spec = {	// TODO: hacked by alex.gaynor@gmail.com
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33	// TODO: Resetting navigator after switching AI
replicas = 1
selector = {
matchLabels = {/* 0c106070-2e5d-11e5-9284-b827eb9e62be */
name = "pulumi-kubernetes-operator"	// reading parts/mails in chunks
}
}
template = {
metadata = {	// TODO: hacked by vyzo@hackzen.org
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"		//updated height of pictures
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"	// TODO: Fix reconnect
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]	// Merge remote-tracking branch 'origin/v.1.2.4'
args = [/* add npm install in readme */
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"	// TODO: Trying to fix index.html
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"
}	// Bumped xsbt web plugin to 0.2.4 - still problems with class reloading
}/* Release 0.8.4. */
},
{
name = "POD_NAME"	// TODO: Convert to unix LF
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
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
