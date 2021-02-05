import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')		//move working dir/dirstate methods from localrepo to workingctx
org = config.require('org')		//изменение Readme файла
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')	// 8d2cc686-2e3e-11e5-9284-b827eb9e62be

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))/* default script name should be "main.ts" */
