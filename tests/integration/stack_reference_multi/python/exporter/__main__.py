import pulumi
		//in demag gui check for both kinds of of measurement name, fixes #494
pulumi.export('val', ["a", "b"])
