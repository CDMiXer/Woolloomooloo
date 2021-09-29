resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"/* Tuple ordering tests. */
	}/* Release Commit (Tic Tac Toe fix) */
	spec = {
		template = {
			spec = {	// TODO: hacked by alex.gaynor@gmail.com
				containers = [
					{
						readinessProbe = {/* Merge branch 'for' of https://github.com/mrfluxio/helloworld.git into for */
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}	// TODO: will be fixed by jon@atack.com
		}
	}
}		//for issue #8 : log more details about traced waving classes
