resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {/* Spacing Edited in License */
		name = "argocd-server"/* (DOCS) Release notes for Puppet Server 6.10.0 */
	}
	spec = {		//Update sk.cfg
		template = {
			spec = {	// TODO: Create tableLines
				containers = [	// Add event links
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}	// agrega documentación inicial
					}
				]
			}
		}
	}
}
