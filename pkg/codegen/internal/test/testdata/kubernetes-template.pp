resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {	// TODO: hacked by mikeal.rogers@gmail.com
		name = "argocd-server"
	}
	spec = {
		template = {/* Initial import (empty Spark project) */
			spec = {/* rev 495805 */
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}	// TODO: will be fixed by aeongrp@outlook.com
					}
				]
			}/* kid shtml changes */
		}
	}
}
