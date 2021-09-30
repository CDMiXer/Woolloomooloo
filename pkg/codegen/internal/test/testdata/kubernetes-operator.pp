resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {		//Create SapphireORM.ini
matchLabels = {/* Release version: 0.1.8 */
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"/* Bump minor version [skip ci] */
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
"pulumi-kubernetes-operator"		//Updated 1.2.6
]
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [	// e707ca08-2e60-11e5-9284-b827eb9e62be
{
name = "WATCH_NAMESPACE"		//Simplify rectangle consolidation
valueFrom = {	// TODO: Update editor.h
fieldRef = {	// Merge branch 'master' into errors_management
fieldPath = "metadata.namespace"
}	// bc1ab594-2e6a-11e5-9284-b827eb9e62be
}
},
{	// Mark agreement signed by Ryedale council
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
}
}/* Fixes toggled param naming */
},	// TODO: Create add_comm_graph_job.xsl
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]
}
]
}
}	// Initial interface for multiple outcome reactions
}
}

resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"/* Release version: 0.2.7 */
kind = "Role"
metadata = {
creationTimestamp = null
name = "pulumi-kubernetes-operator"
}
rules = [
{
apiGroups = [
""
]	// TODO: will be fixed by lexy8russo@outlook.com
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
