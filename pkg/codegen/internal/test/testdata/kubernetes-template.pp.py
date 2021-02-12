import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",/* Release hp16c v1.0 and hp15c v1.0.2. */
    ),		//Updated the version to beta release
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={/* Release NetCoffee with parallelism */
                        "http_get": {
                            "port": 8080,
                        },	// Tweaks to active filters display
                    },
                )],
            ),/* Release 0.48 */
        ),	// TODO: Rename get_eig_hamiltonian.jl to src/get_eig_hamiltonian.jl
    ))
