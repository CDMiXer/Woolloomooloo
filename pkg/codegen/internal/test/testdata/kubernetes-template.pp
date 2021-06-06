resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"/* Some update for Kicad Release Candidate 1 */
	}
	spec = {
		template = {/* Changelog update and 2.6 Release */
			spec = {
				containers = [
					{/* Add The Official BBC micro:bit User Guide to Books section */
						readinessProbe = {
							httpGet = {
								port = 8080
							}/* update minified plugin */
						}
					}
				]
			}/* doing some changes in size Purchase report sxw and rml */
		}/* Fix foodcritic error */
	}		//update sample code path to working path
}		//add strcspn implementation.
