resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {	// TODO: Adding in the apparmor profile
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}
					}/* Release of eeacms/www:18.4.2 */
				]
			}		//send init events
		}
	}
}
