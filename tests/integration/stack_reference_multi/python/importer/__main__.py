import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')	// TODO: Update M003.md
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')	// DynamoDB added as a supported database

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
