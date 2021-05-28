import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')		//Rewrote network proxy to byte oriented protocol
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))	// Merge "Handle not found in check for disk availability"
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
