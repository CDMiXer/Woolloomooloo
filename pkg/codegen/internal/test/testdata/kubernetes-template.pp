resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {/* Merge "update constraint for oslo.rootwrap to new release 6.0.0" */
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {/* Fix hierarchy, better flow */
								port = 8080
							}
						}
					}	// TODO: convert files fro ASCII to UTF8 (zombielei's patch, part 2)
				]	// TODO: Node based ops for Unicorn
			}
		}
	}
}
