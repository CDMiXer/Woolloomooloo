import pulumi
		//Fixed error in __all__ declaration
config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))		//upped default bootstrap timeout.
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
