resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {/* Updated JSON generator */
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080		//Merge branch 'master' into pilot-schools-about
							}
						}	// TODO: hacked by arajasek94@gmail.com
					}
				]	// work in progress improvements.
			}
		}
	}
}
