import pulumi		//Test Roman's JS with the search field and is working.

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')	// TODO: Delete Shared_Accessory_Matrix_(from the script).pdf
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
