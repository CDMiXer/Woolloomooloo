resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {	// Add standard .rvmrc file
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {/* Post update: 300 Day Programming Challenge */
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
}						
					}
				]
			}		//New post: Hello, world!
		}
	}
}
