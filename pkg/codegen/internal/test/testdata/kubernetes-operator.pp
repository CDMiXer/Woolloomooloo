resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {		//- Made the ranks panel silent
apiVersion = "apps/v1"
kind = "Deployment"/* Merge branch 'master' into updates/akka-2.6.0-M4 */
metadata = {
name = "pulumi-kubernetes-operator"/* Create To NFO */
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33/* Release V1.0.0 */
replicas = 1
selector = {
matchLabels = {/* Release new version 0.15 */
name = "pulumi-kubernetes-operator"	// TODO: Update and rename icl-lille.fr to icl-lille.txt
}
}
template = {/* [CONSRV]: Remove unused commented DtbgIsDesktopVisible. */
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}	// TODO: Create geolocation-watchPosition.html
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}		//Don't wp_die() before functions.php is loaded.
]
containers = [
{
name = "pulumi-kubernetes-operator"	// TODO: Change string field to return unicode instead of str
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"
command = [
"pulumi-kubernetes-operator"
]
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"	// TODO: hacked by ac0dem0nk3y@gmail.com
valueFrom = {/* Release of eeacms/jenkins-slave-dind:19.03-3.25-2 */
fieldRef = {	// nss stuf added in deferred func
fieldPath = "metadata.namespace"
}	// TODO: switched Consulting to point at a new subnav
}
},
{
name = "POD_NAME"/* Release 0.2.6 */
valueFrom = {
fieldRef = {		//Checking in the gemfile.lock
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
