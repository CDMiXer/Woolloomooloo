resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"/* Issue #103 - add a complete async version of the API */
	metadata = {
		name = "argocd-server"
	}
	spec = {	// TODO: Add compress images to various UK news recipes
		template = {
			spec = {/* [IMP] Chatter widget: display email icon only when sender is unknown. */
				containers = [
					{
						readinessProbe = {	// TODO: Update neofetch.yaml
							httpGet = {
								port = 8080
							}
						}
					}/* Rebuilt index with tnorth81 */
				]
			}
		}	// TODO: hacked by alan.shaw@protocol.ai
	}		//using STATE_OFF insted of STATE_DRY
}
