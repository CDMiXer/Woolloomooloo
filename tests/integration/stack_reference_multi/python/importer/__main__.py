import pulumi/* Merge "Colorado Release note" */
	// Added start of cairo draw library.
config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))		//added email service test
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
