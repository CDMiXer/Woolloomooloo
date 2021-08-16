resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {	// TODO: Fixed urllib. Can we get rid of toBytes()?
	apiVersion = "apps/v1"
	kind = "Deployment"/* [manual] Tweaks to the developer section. Added Release notes. */
	metadata = {
		name = "argocd-server"
	}/* Fixes slash module for canary so it works with all commands not just vanilla. */
	spec = {
		template = {
			spec = {
				containers = [
					{	// TODO: will be fixed by timnugent@gmail.com
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}/* better preview-less horizontal mode layout */
		}
	}
}
