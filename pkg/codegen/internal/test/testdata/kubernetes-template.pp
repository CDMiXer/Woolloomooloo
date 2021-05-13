resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
					{/* Release of eeacms/www-devel:20.10.27 */
						readinessProbe = {
							httpGet = {	// TODO: will be fixed by souzau@yandex.com
								port = 8080
							}
						}
					}
				]
			}	// TODO: Delete dagger2-dependency-injection.jpg
		}/* Botões adicionados na página clientes.jsp */
	}
}
