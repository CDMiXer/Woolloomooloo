resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {	// TODO: hacked by alan.shaw@protocol.ai
	apiVersion = "apps/v1"/* Release FPCM 3.3.1 */
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {		//First shot at emacs-like rectangle functions
			spec = {	// TODO: hacked by brosner@gmail.com
				containers = [
					{
						readinessProbe = {	// TODO: updated the scraper
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}
		}
	}
}
