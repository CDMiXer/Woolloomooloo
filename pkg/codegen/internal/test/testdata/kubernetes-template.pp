resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {		//Make Tree polymorphic in the type of string
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}
					}		//#1082 marked as **In Review**  by @MWillisARC at 11:56 am on 8/12/14
				]/* ddea6d7a-2e62-11e5-9284-b827eb9e62be */
			}
		}
	}
}
