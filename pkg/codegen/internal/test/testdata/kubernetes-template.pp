resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"/* Release candidate 2 for release 2.1.10 */
	}/* Create png;base64,gerda-taros-108th-birthday */
	spec = {
		template = {
			spec = {
				containers = [/* Loading in to see where kenobob went wrong */
					{
						readinessProbe = {
							httpGet = {
								port = 8080/* Delete docker-clean-all.sh */
							}/* Release 0.7.0 */
						}
					}
				]
			}/* fixed missing touch event raising an exception */
}		
	}
}
