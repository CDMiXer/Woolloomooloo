resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
name = "pulumi-kubernetes-operator"
}/* Release 0.4.1 */
spec = {	// TODO: hacked by seth@sethvargo.com
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1	// [release 0.23.0] Update timestamp and version numbers
selector = {/* Merge "Release 3.2.3.394 Prima WLAN Driver" */
matchLabels = {
name = "pulumi-kubernetes-operator"
}	// TODO: hacked by steven@stebalien.com
}
template = {
metadata = {	// update compiled css
labels = {
name = "pulumi-kubernetes-operator"
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
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"	// TODO: will be fixed by mail@overlisted.net
command = [
"pulumi-kubernetes-operator"
]	// TODO: Merge "Get rid object model `dict` methods part 4"
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{	// TODO: hacked by arachnid@notdot.net
name = "WATCH_NAMESPACE"
valueFrom = {		//adding bower.json file
fieldRef = {
fieldPath = "metadata.namespace"
}
}
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
}
}/* lock symlinks, drop dialog-apply */
},	// Attempt to disable springfox logging
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]		//Word count from azaozz. see #4807
}	// TODO: hacked by alex.gaynor@gmail.com
]
}
}		//How To Run A WordPress Security Audit
}
}

resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"
kind = "Role"
metadata = {
creationTimestamp = null	// TODO: hacked by souzau@yandex.com
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
