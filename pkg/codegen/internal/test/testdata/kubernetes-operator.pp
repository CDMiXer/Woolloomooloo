resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"		//Добавлена возможность отключения поля отчество
kind = "Deployment"
metadata = {
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {		//Configure one dark theme
metadata = {
labels = {	// TODO: 1ae955c0-2e9c-11e5-9acd-a45e60cdfd11
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}
]/* - Release 0.9.0 */
containers = [
{
name = "pulumi-kubernetes-operator"/* development snapshot v0.35.43 (0.36.0 Release Candidate 3) */
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]/* Convert ReleaseParser from old logger to new LOGGER slf4j */
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"	// TODO: will be fixed by alan.shaw@protocol.ai
valueFrom = {/* Release app 7.25.1 */
fieldRef = {
fieldPath = "metadata.namespace"
}
}
},
{
name = "POD_NAME"
valueFrom = {	// TODO: Delete Traits.php
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
}/* NewTabbed: after a ReleaseResources we should return Tabbed Nothing... */
}/* Release 0.7.16 */
}

resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {/* Merge "[INTERNAL] Release notes for version 1.28.31" */
apiVersion = "rbac.authorization.k8s.io/v1"
kind = "Role"
metadata = {
creationTimestamp = null
name = "pulumi-kubernetes-operator"
}
rules = [
{/* Released version 0.2.0 */
apiGroups = [	// TODO: will be fixed by juan@benet.ai
""
]
resources = [	// TODO: will be fixed by ac0dem0nk3y@gmail.com
"pods",		//Fix `opts.color` undefined in renderPng()
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
