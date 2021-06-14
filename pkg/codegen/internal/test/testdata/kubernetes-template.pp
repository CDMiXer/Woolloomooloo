resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"/* update bower json to v1.1.0 */
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [		//Fixed configure and Makefiles
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}/* Release Notes for v02-15-04 */
					}
				]
			}
		}
	}
}
